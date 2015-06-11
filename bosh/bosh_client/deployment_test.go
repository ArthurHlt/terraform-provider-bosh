package bosh_client

import (
	"log"
	"testing"

//	"github.com/kr/pretty"
)

func TestExecuteDeployment(t *testing.T) {

	var (
		err error
		
		cloudConfig string
		
		cloudConfigManifest *Manifest
		deploymentManifest *Manifest
		
		deployments []*Deployment
	)
	
	d := GetInitializedDirector(t)
	
	// Update cloud config
	
	cloudConfigManifest, err = d.NewManifest(cloudConfig1)		
	err = d.UpdateCloudConfig(cloudConfigManifest)
	if err != nil {
		log.Printf("[FAIL] Error uploading cloud-config: %s", err.Error())
		t.FailNow()
	}
	
	cloudConfig, err = d.GetCloudConfig()
	if err != nil {
		log.Printf("[FAIL] Error retrieving cloud-config: %s", err.Error())
		t.FailNow()		
	}
	
	assertPattern(t, cloudConfig, "compilation:")
	assertPatternFalse(t, cloudConfig, "- range: 10.244.8.0/30")
	
	cloudConfigManifest, err = d.NewManifest(cloudConfig2)		
	err = d.UpdateCloudConfig(cloudConfigManifest)
	if err != nil {
		log.Printf("[FAIL] Error uploading cloud-config: %s", err.Error())
		t.FailNow()
	}
	
	cloudConfig, err = d.GetCloudConfig()
	if err != nil {
		log.Printf("[FAIL] Error retrieving cloud-config: %s", err.Error())
		t.FailNow()		
	}
	
	assertPattern(t, cloudConfig, "- range: 10.244.8.0/30")

	// Deploy

	deploymentManifest, err = d.NewManifest(dockerManifest)
	if err != nil {
		log.Printf("[FAIL] Error creating manifest: %s", err.Error())
		t.FailNow()
	}

	deploymentManifest.Properties["testVar1"] = "value1"
	deploymentManifest.Properties["testVar2"] = "value2"

	err = d.Deploy(deploymentManifest)
	if err != nil {
		log.Printf("[FAIL] Error deploying manifest: %s", err.Error())
		t.FailNow()		
	}

	// Verify deployment
	
	deployments, err = d.ListDeployments()
	if err != nil {
		log.Printf("[FAIL] Error retrieving list of deployments: %s", err.Error())
		t.FailNow()
	}
	if len(deployments) != 1 || deployments[0].Name != "docker" {		
		log.Printf("[FAIL] Expected 1 deployments with name 'docker' but got %d deployments with name '%s'.", len(deployments), deployments[0].Name)
		t.FailNow()
	}
	deployment := deployments[0]
	if len(deployment.Releases) != 1 || deployment.Releases[0] != "docker/13" {
		log.Printf("[FAIL] Expected 1 'docker/13' releases to have been deployed but got %d '%s' releases.", len(deployment.Releases), deployment.Releases[0])
		t.FailNow()		
	}
}

func TestDeleteDeployment(t *testing.T) {

	d := GetDirector(t)
	
	// Delete deployment
	
	deployments, err := d.ListDeployments()
	if err != nil {
		log.Printf("[FAIL] Error retrieving list of deployments: %s", err.Error())
		t.FailNow()
	}
	
	deployment := deployments[0]
	log.Printf("[DEBUG] Deleting deployment '%s'.", deployment.Name)
	err = deployment.delete()
	if err != nil {
		log.Printf("[FAIL] Error deleting deployments: %s", err.Error())
		t.FailNow()
	}

	// Verify deletion of deployment

	deployments, err = d.ListDeployments()
	if err != nil {
		log.Printf("[FAIL] Error retrieving list of deployments: %s", err.Error())
		t.FailNow()
	}
	if len(deployments) != 0 {
		log.Printf("[FAIL] Expected 0 deployments but got %d deployments.", len(deployments))
		t.FailNow()		
	}
}

const cloudConfig1 = `---
compilation:
  workers: 1
  network: default
  reuse_compilation_vms: true
  cloud_properties: {}

networks:
- name: default
  subnets: []

resource_pools:
  - name: default
    network: default
    stemcell:
      name: bosh-warden-boshlite-ubuntu-trusty-go_agent
      version: {{ .LatestStemcellVersion "bosh-warden-boshlite-ubuntu-trusty-go_agent" }}
    cloud_properties: {}
`

const cloudConfig2 = `---
compilation:
  workers: 1
  network: default
  reuse_compilation_vms: true
  cloud_properties: {}

networks:
  - name: default
    subnets:
      # network with static ip used for web
      - range: 10.244.8.0/30
        gateway: 10.244.8.1
        # reserved: [10.244.8.1]
        static: [10.244.8.2]
        cloud_properties: {}
      # networks for dynamic ips (db, workers, compilation vms)
      - range: 10.244.8.4/30
        gateway: 10.244.8.5
        # reserved: [10.244.8.5]
        cloud_properties: {}
      - range: 10.244.8.8/30
        gateway: 10.244.8.9
        # reserved: [10.244.8.9]
        cloud_properties: {}
      - range: 10.244.8.12/30
        gateway: 10.244.8.13
        # reserved: [10.244.8.13]
        cloud_properties: {}
      - range: 10.244.8.16/30
        gateway: 10.244.8.17
        # reserved: [10.244.8.17]
        cloud_properties: {}
      - range: 10.244.8.20/30
        gateway: 10.244.8.21
        # reserved: [10.244.8.21]
        cloud_properties: {}

resource_pools:
  - name: default
    network: default
    stemcell:
      name: bosh-warden-boshlite-ubuntu-trusty-go_agent
      version: {{ .LatestStemcellVersion "bosh-warden-boshlite-ubuntu-trusty-go_agent" }}
    cloud_properties: {}
`

const dockerManifest = `---
name: docker
director_uuid: {{ .DirectorUUID }}

releases:
 - name: docker
   version: {{ .LatestReleaseVersion "docker" }}

update:
  canaries: 0
  canary_watch_time: 30000-240000
  update_watch_time:  30000-240000
  max_in_flight: 32
  serial: false

jobs:
  - name: docker
    templates:
      - name: docker
      - name: containers
    instances: 1
    resource_pool: default
    persistent_disk: 65536
    networks:
      - name: default
    properties:
      containers:
        - name: redis
          image: "redis"
          command: "--dir /var/lib/redis/ --appendonly yes"
          bind_ports:
            - "6379:6379"
          bind_volumes:
            - "/var/lib/redis"
          entrypoint: "redis-server"
          memory: "256m"
          cpu_shares: 1
          env_vars:
            - "TEST_VAR1={{ .P "testVar1" }}"
            - "TEST_VAR2={{ .P "testVar2" }}"
`
