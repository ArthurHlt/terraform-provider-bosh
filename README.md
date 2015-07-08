# Bosh Provider for Terraform [![Build Status](https://travis-ci.org/mevansam/terraform-provider-bosh.svg?branch=master)](https://travis-ci.org/mevansam/terraform-provider-bosh)

The [Bosh](http://bosh.io/) provider for [Terraform](https://terraform.io/) integrates the Bosh deployment and operations toolset with a Terraform'ed cloud. All work for this provider is tracked using [PivotalTracker](https://www.pivotaltracker.com/projects/1359482). 

To contribute:

1. Fork the project and make your changes and submit a pull request.
2. Request to be added to the list of contributors to this Github repository and the PivotalTracker project.

## Provider

The Bosh provider will target the given endpoint if it is available. If not the "bosh-deploy" resource can be used to build the micro-bosh deployment to be targeted.

```
provider "bosh" {
    target = "#.#.#.#"  # Publicly addressable name or IP
    username = "admin"
    password = "*****"
}
```

## Terraform Resources

The Bosh Terraform resources describe the elements of the manifest used by Bosh to build the infrastructure. These resources can be considered to be logical representations of physical resources managed by Bosh. The following diagram outlines the relationships between the various elements of a manifest. The decomposition of the manifest elements into Terraform resources will also make it easier to adapt to the evolving structure of the Bosh manifest, as outlined in the [bosh notes](https://github.com/cloudfoundry/bosh-notes).

![Image of Bosh Manifest Elements]
(docs/images/Bosh-Manifest-Elements.png)

### Assets

*bosh_stemcell* and *bosh_release* identify the software assets required to build the environment. 

#### "bosh_stemcell"

The *bosh_stemcell* resource describes a base operating system image that can be referenced by bosh resources.

```
resource "bosh_stemcell" "ubuntu" {

    name = "bosh-openstack-kvm-ubuntu-trusty-go_agent"
    version = "2978"
    url = "https://bosh.io/d/stemcells/bosh-openstack-kvm-ubuntu-trusty-go_agent?v=2978"
    sha1 = "42e08d492dafd46226676be59ee3e6e8a0da618b" 
}
```

The *sha1* digest is required for stemcells referred to by the *bosh_microbosh* resource. It is optional for all other resources that refer to a stemcell. The *name* and *version* attribute values must match that of the actual stemcell and will be validated once the stemcell has been uploaded to bosh.


#### "bosh_release"

The *bosh_release* resource describes software packages and configuration used by a bosh deployment.

```
resource "bosh_release" "bosh" {

    name = "bosh"
    version = "175"
    url = "https://bosh.io/d/github.com/cloudfoundry/bosh?v=175"
    sha1 = "be849a22b9034fc7a682d72ee1aaa84aebb2b8e5 "
}
```

The *sha1* digest is required for releases referred to by the *bosh_microbosh* resource. It is optional for all other resources that refer to a release. The *name* and *version* attribute values must match that of the actual release and will be validated once the release has been uploaded to bosh.

### Infrastructure

*bosh_availability_zone*, *bosh_network*, *bosh_disk* and *bosh_resource* identify infrastructure used by bosh jobs to create IaaS resources during deployment. The phyiscal infrastructure specific attributes of these logical resources are specified in the *bosh_cloud_config* resource. A *bosh_job* resource should only reference the logical infrastructure resources thus maintaining cloud portability.

#### "bosh_availability_zone"

The *bosh_availability_zone* describes an availability zone or placement pool as defined (here)[https://github.com/cloudfoundry/bosh-notes/blob/master/availability-zones.md]. 

```
resource "bosh_availability_zone" "az1" {
	name = "az1"
}
resource "bosh_availability_zone" "az2" {
	name = "az2"
}
```

#### "bosh_network"

The *bosh_network* resource describes a network resource use by a bosh deployment. It is used to build the network section of the bosh cloud config manifest and it's attributes closely map to the [documentation published at bosh.io](http://bosh.io/docs/networks.html). The network resource also acts as simple IPAM resource for Jobs requiring static IPs. The IP allocations are persisted in the Terraform state file.

```
resource "bosh_network" "public-network" {

    name = "public"
    type = "vip"
}
resource "bosh_network" "infra-network" {

    name = "infra"
    type = "manual"
        
    static_ip_block {
        name = "proxy"
        num_ips = 2
    }
}
resource "bosh_network" "application-network" {

    name = "application"
    type = "manual"

    num_reserved_ips = 100
        
    static_ip_block {
        name = "proxy"
        num_ips = 2
    }
    static_ip_block {
        name = "router"
        num_ips = 2
    }
}
```

Computed attributes:

* id - id used to reference a bosh resource network instance
* vip.# - number of vips allocated to the network
* vip.[*N*] - the Nth vip 
* static.[*STATIC_IP_BLOCK_NAME*].count - number of allocated static ips
* static.[*STATIC_IP_BLOCK_NAME*].[*N*].ip - the Nth static ip 

If more than one placement pool (or availability zones) is available in the cloud then the Nth static IP for an instance in each availability zone will be unique as the total number of static IPs will be the number of IPs requested times the number of pools.


#### "bosh_disk""

The *bosh_persistent_disk* resource describes a named disk type for a bosh deployment. It is used to build the disk pool section of a bosh cloud config manifest and the resource attributes closely map to the [documentation published at bosh.io](http://bosh.io/docs/persistent-disks.html).

```
resource "bosh_disk" "fast" {
    name = "fast-disk"
    disk_size = 512
}
resource "bosh_disk" "standard-medium" {
    name = "standard-medium-disk"
    disk_size = 512
}
resource "bosh_disk" "standard-large" {
    name = "standard-large-disk"
    disk_size = 1024
}
```

Computed attributes:

* id - ID used to reference a bosh resource instance.

#### "bosh_compilation"

The *bosh_compilation* resource describes the attributes of compilation jobs.

```
resource "bosh_compilation" "compilation" {
    workers = 8
    reuse_compilation_vms = true
    network = "${bosh_network.infra-network.name}"
}
```

Computed attributes:

* id - ID used to reference a bosh resource instance.

#### "bosh_resource"

The *bosh_resource* resource describes the IaaS resource for a bosh deployment job. It is used to build the resource pool section of a bosh cloud config manifest and the resource attributes closely map to the [documentation published at bosh.io](http://bosh.io/docs). In order to reference this resource from a *bosh_job* resource it's infratructure specific attributes must have been updated by a referring *bosh_cloud_config* resource.

```
resource "bosh_resource" "small" {
    name = "cf_small"
}
resource "bosh_resource" "medium" {
    name = "cf_medium"
}
resource "bosh_resource" "large" {
    name = "cf_large"
}
```

Computed attributes:

* id - id used to reference a bosh resource instance

#### "bosh_job"

The *bosh_job* resource describes a job or instance of a runable process.

```
resource "bosh_job" "postgres" {
	
	instances = 1
	resource_pool {
		ref = "${bosh_resource.large.id}"
	}
	network {
		ref = "${bosh_network.application-network.id}"
	}
}
resource "bosh_job" "uaa" {

	instances = 2
	resource_pool {
		ref = "${bosh_resource.medium.id}"
	}
	network {
		ref = "${bosh_network.application-network.id}"
	}
}
```

### Cloud

#### "bosh_cloud_config"

The *bosh_cloud_config* composes the *bosh_network*, *bosh_disk* and *bosh_resource* resources to build cloud config manifest that can by pushed to the bosh director before executing a deployment.

```
resource "bosh_cloud_config" "openstack_dev" {

    availability_zone {
        ref = "${bosh_availability_zone.az1.id}"
        
        cloud_property {
            name = "availability_zone"
            value = "USEAST_AZ1"
        }
    }
    availability_zone {
        ref = "${bosh_availability_zone.az2.id}"

        cloud_property {
            name = "availability_zone"
            value = "USEAST_AZ2"
        }
    }

    network {
        ref = "${bosh_network.public-network.id}"
        
        vip = [
            "${openstack_compute_floatingip_v2.vip1.address}",
            "${openstack_compute_floatingip_v2.vip2.address}",
            "${openstack_compute_floatingip_v2.vip3.address}",
            "${openstack_compute_floatingip_v2.vip4.address}"
        ]
    }
    network {
        ref = "${bosh_network.infra-network.id}"
        
        subnet {
            cidr = "${openstack_networking_subnet_v2.bosh_apps_subnet.cidr}"
            gateway = "${openstack_networking_subnet_v2.bosh_apps_subnet.gateway_ip}"
            
            availability_zone = "self.availability_zone.0.name"
            
            cloud_property {
                name = "net_id"
                value = "${openstack_networking_subnet_v2.bosh_infra_network.id}"
            }
        }
    }
    network {
        ref = "${bosh_network.application-network.id}"
        
        subnet {
            cidr = "${openstack_networking_subnet_v2.bosh_infra_subnet.cidr}"
            gateway = "${openstack_networking_subnet_v2.bosh_infra_subnet.gateway_ip}"
            
            cloud_property {
                name = "net_id"
                value = "${openstack_networking_subnet_v2.bosh_apps_network.id}"
            }
        }
    }
    
	resource {
    	ref = "${bosh_resource.small.id}"
        
        stemcell {
        	name = ${bosh_stemcell.ubuntu.name}"
            varsion = ${bosh_stemcell.ubuntu.version}"
        }
        
        cloud_property {
        	name = "instance_type"
            value = "m1.small"
        }
	}
	resource {
    	ref = "${bosh_resource.medium.id}"
        
        stemcell {
        	name = ${bosh_stemcell.ubuntu.name}"
            varsion = ${bosh_stemcell.ubuntu.version}"
        }
    
        cloud_property {
        	name = "instance_type"
            value = "m1.medium"
        }
	}
	resource {
    	ref = "${bosh_resource.large.id}"
        
        stemcell {
        	name = ${bosh_stemcell.ubuntu.name}"
            varsion = ${bosh_stemcell.ubuntu.version}"
        }
    
        cloud_property {
        	name = "instance_type"
            value = "m1.large"
        }
	}
    
    disk {
    	ref = "${bosh_disk.fast-disk.id}"

        cloud_property {
        	name = "type"
            value = "gp2"
        }
    }
    disk {
    	ref = "${bosh_disk.standard-disk-medium.id}"

        cloud_property {
        	name = "type"
            value = "standard"
        }
    }
    disk {
    	ref = "${bosh_disk.standard-disk-large.id}"

        cloud_property {
        	name = "type"
            value = "standard"
        }
    }
    
    compilation {
        ref = "${bosh_compilation.compilation.id}"
        
        cloud_property {
        	name = ""instance_type"
            value = "m1.medium"
        }
        cloud_property {
        	name = ""availability_zone"
            value = "USEAST_AZ1"
        }
    }
}
```

### Deployment

#### "bosh_deployment"

```
resource "bosh_deployment" "docker" {

	job {
    	name = "docker"
        templates = [
        	"docker",
            "containers"
        ]
        resource_pool = "${bosh_resource.infra_medium.id}"
        disk_pool = "${bosh_disk.fast_disk.id}"
    }
}
```

### "bosh_manifest_template"

The *bosh_manifest_template* resource identifies a templated manifest file that can be used by the *bosh_deployment* resource. This is an alternate approach to building a bosh deployment compared to using the more granular resource constructs described previously. 

```
resource "bosh_manifest_template" "openstack_cloud_mf" {

    template = "/path/to/template"
}
resource "bosh_cloud_config" "openstack_cloud" {

    manifest = "${bosh_manifest_template.openstack_cloud_mf.id}"
}
resource "bosh_deployment" "docker" {

    name = "docker"    

    deployment_manifest = "${bosh_manifest_template.docker_deployment_mf.id}"
    cloud_config_manifest = "{bosh_manifest_template.openstack_cloud_mf.id}"
}
```

### "bosh_microbosh"

Deploys Microbosh to the specified IaaS using *[bosh-init](https://github.com/cloudfoundry/bosh-init)*. If this resource is specified and the ```target``` given in the provider configuration is not pingable, then this resource will attempt to create the Microbosh instance.

```
resource "bosh_microbosh" {

    name = "my_terraformed_microbosh"
    
    binary = true  # set up microbosh in a binary configuration for HA
    
    release {
        url = "${bosh_release.bosh.url}"
        sha1 = "${bosh_release.bosh.sha1}"
    }
    
    stemcell {
        url = "${bosh_stemcell.ubuntu.url}"
        sha1 = "${bosh_stemcell.ubuntu.sha1}"
    }

    ###############################
    # Bosh IaaS CPI configuration #

    openstack {
        
        username = "os_user"
        password = "os_password"
        tenant = "bosh_tenant"
        auth_url = "https://my-openstack.com:5000/v2.0"
        region = "east"
    }

    # OR
    vsphere {
    }

    # OR
    aws {
    }

    # OR
    google {
    }
}
```

### "bosh_deployment"

## Running Tests

The Bosh integration tests are run against a local Bosh Director instance. This can be launched using the *Vagrantfile* located in the project root.

The vagrant-bosh plugin needs to be installed so that Vagrant can provision the Bosh bits. 

To install the plugin:
```
$ vagrant plugin install vagrant-bosh
```

And then run:
```
$ vagrant up
```

Validate that the machine was provisioned correctly by running:
```
$ bosh --target 127.0.0.1 --user admin --password admin status
```

Which should output:
```
Config
             /Users/msamaratunga/.bosh_config

Director
  Name       Bosh Lite Director
  URL        https://127.0.0.1:25555
  Version    1.2977.0 (00000000)
  User       admin
  UUID       7d926be9-d123-4155-987e-65e6c2879f98
  CPI        cpi
  dns        disabled
  compiled_package_cache enabled (provider: local)
  snapshots  disabled

Deployment
  not set
```

To run the tests execute the following shell commands from the bosh provider directory. These integration tests will take a long time to complete as they will upload stemcells / releases and run deployments mutliple times. They need access to the internet to complete successfully and a fast network connection will speed up the tests.

```
$ export TF_ACC=1
$ export BOSH_TARGET=127.0.0.1
$ export BOSH_USER=admin
$ export BOSH_PASSWORD=password
$ go test -v
```

## License and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Author | Email | Company
-------|-------|--------
Mevan Samaratunga | msamaratunga@pivotal.io | [Pivotal](http://www.pivotal.io)
