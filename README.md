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

### "bosh_stemcell"

Describes a stemcell that can be referenced by bosh deployments. 

```
resource "bosh_stemcell" "ubuntu" {

    url = "https://bosh.io/d/stemcells/bosh-openstack-kvm-ubuntu-trusty-go_agent?v=2978"
    
    sha1 = "42e08d492dafd46226676be59ee3e6e8a0da618b" 
}
```

The *sha1* digest is required for stemcells referred to by the *bosh_microbosh* resource. It is optional for all other resources that refer to a stemcell, but if provided it will be validated by downloading it to a temp folder.

Computed attributes:

* name

### "bosh_release"

Describes a bosh release that it can be referenced by bosh deployments.

```
resource "bosh_release" "bosh" {

    url = "https://bosh.io/d/github.com/cloudfoundry/bosh?v=169"
    
    sha1 = "ec361150584094951273f1088344e2d4b2ebeb9f"
}
```

The *sha1* digest is required for releases referred to by the *bosh_microbosh* resource. It is optional for all other resources that refer to a release, but if provided it will be validated by downloading it to a temp folder.

Computed attributes:

* name

### "bosh_cloud_config"


```
resource "bosh_cloud_config" "openstack_dev" {

	availability_zones = [ "AZ1", "AZ2" ]

    network {
        name = "public"
        type = "vip"
    }
	network {
		name = "infra-network"
	    type = "manual"

	    cidr = "${openstack_networking_subnet_v2.bosh_infra_subnet.cidr}"
    	gateway = "${openstack_networking_subnet_v2.bosh_infra_subnet.gateway_ip}"
        num_reserved_ips = 100
        
        static_ip_block {
        	name = "proxy"
            num_ips = 2
        }
    	
        cloud_property {
        	name = "net_id"
            value = "${openstack_networking_subnet_v2.bosh_infra_network.id}"
        }
    }
	network {
		name = "application-network"
	    type = "manual"

	    cidr = "${openstack_networking_subnet_v2.bosh_apps_subnet.cidr}"
    	gateway = "${openstack_networking_subnet_v2.bosh_apps_subnet.gateway_ip}"
        num_reserved_ips = 100
        
        static_ip_block {
        	name = "proxy"
            ips_per_az = 2
        }
        static_ip_block {
        	name = "router"
            ips_per_az = 2
        }
    	
        cloud_property {
        	name = "net_id"
            value = "${openstack_networking_subnet_v2.bosh_apps_network.id}"
        }
    }
    
	resource {    
    	name = ${bosh_resource.infra_medium.name}
        
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
    	name = ${bosh_resource.infra_large.name}
        network = ${bosh_resource.infra_large.network}
        
        stemcell {
        	name = ${bosh_stemcell.ubuntu.name}"
            varsion = ${bosh_stemcell.ubuntu.version}"
        }
    
        cloud_property {
        	name = "instance_type"
            value = "m1.large"
        }
	}
	resource {    
    	name = ${bosh_resource.app_medium.name}
        network = ${bosh_resource.app_medium.network}
        
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
    	name = ${bosh_resource.app_large.name}
        network = ${bosh_resource.app_large.network}
        
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
    }
    disk {
    }
    
    compilation {
    }
}
```

Computed attributes:

- networks.[*NETWORK_NAME*].vip.count - number of vips allocated to the network
- networks.[*NETWORK_NAME*].vip.[*N*].ip - the Nth vip 
- networks.[*NETWORK_NAME*].static.[*STATIC_IP_BLOCK_NAME*].count - number of allocated static ips
- networks.[*NETWORK_NAME*].static.[*STATIC_IP_BLOCK_NAME*].[*N*].ip - the Nth static ip 


### "bosh_deployment"

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

### "bosh_microbosh"

Deploys Microbosh to the specified IaaS using *[bosh-init](https://github.com/cloudfoundry/bosh-init)*. If this resource is specified and the ```target``` given in the provider configuration is not pingable, then this resource will attempt to create the Microbosh instance.

```
resource "bosh_microbosh" {

    name = "my_terraformed_microbosh"
    
    binary = true  # set up microbosh in a binary configuration for HA
    
    release {
        url = "${bosh_release.bosh.url}"
        url = "${bosh_release.bosh.sha1}"
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
