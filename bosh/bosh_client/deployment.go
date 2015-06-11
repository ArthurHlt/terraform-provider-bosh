package bosh_client

import (
	"fmt"
)

type Deployment struct {
	
	Name string
	Releases []string
	Stemcells []string
	
	CloudConfig string	
	Manifest string	
	
	api *Director
}

func (d *Director) GetCloudConfig() (string, error) {
	
	cc, resp := d.director.GetCloudConfigYaml()
	if resp.IsNotSuccessful() {
		return "", fmt.Errorf("unable to retrieve cloud config properties: %s", resp.Message)
	}
	return cc, nil
}

func (d *Director) UpdateCloudConfig(cc *Manifest) error {
	
	s, err := cc.Process()
	if err != nil {
		return fmt.Errorf("error processing cloud config yaml: %s", err.Error())
	}
		
	resp := d.director.UpdateCloudConfigYaml(s)
	if resp.IsNotSuccessful() {
		return fmt.Errorf("error deploying cloud config yaml: %s", resp.Message)
	}
	return nil
}

func (d *Director) Deploy(m *Manifest) error {
	
	s, err := m.Process()
	if err != nil {
		return fmt.Errorf("error processing manifest: %s", err.Error())
	}
		
	resp := d.director.DeployManifestYaml(s)
	if resp.IsNotSuccessful() {
		return fmt.Errorf("error deploying manifest: %s", resp.Message)
	}
	return nil
}

func (d *Director) ListDeployments() ([]*Deployment, error) {
	
	deployments := []*Deployment{}
	
	deploymentList, resp := d.director.GetDeployments()
	if resp.IsNotSuccessful() {
		return nil, fmt.Errorf("Could not fetch BOSH deployments")
	} else {
		for _, m := range deploymentList {
			
			deployment := Deployment {
				Name: m.Name,
				Releases: []string{},
				Stemcells: []string{},
				CloudConfig: m.CloudConfig,
				api: d,
			}
			
			deployment.Manifest, resp = d.director.GetDeploymentManifestYaml(m.Name)
			if resp.IsNotSuccessful() {
				return nil, fmt.Errorf("error retrieving manifest for deployment '%s': %s", m.Name, resp.Message) 
			} 
			
			for _, r := range m.Releases {
				deployment.Releases = append(deployment.Releases, fmt.Sprintf("%s/%s", r.Name, r.Version))
			}
			for _, v := range m.Releases {
				deployment.Stemcells = append(deployment.Stemcells, fmt.Sprintf("%s/%s", v.Name, v.Version))				
			}
			
			deployments = append(deployments, &deployment)
		}
	}
	
	return deployments, nil
}

func (m *Deployment) delete() error {
	
	resp := m.api.director.DeleteDeployment(m.Name)
	if resp.IsNotSuccessful() {
		return fmt.Errorf("error deleting deployment '%s'", m.Name)
	}
	return nil
}
