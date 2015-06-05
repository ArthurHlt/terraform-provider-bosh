# Bosh Provider for Terraform [![Build Status](https://travis-ci.org/mevansam/terraform-provider-bosh.svg?branch=master)](https://travis-ci.org/mevansam/terraform-provider-bosh)

The [Bosh](http://bosh.io/) provider for [Terraform](https://terraform.io/) provides seamless integration of the Bosh deployment and operations toolset with a Terraform'ed cloud. All work for this provider is tracked using [PivotalTracker](https://www.pivotaltracker.com/projects/1359482). 

To contribute:

1. Fork the project and make your changes and submite a pull request.
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

    stemcell_name = "bosh-stemcell-2978-openstack-kvm-ubuntu-trusty-go_agent.tgz"
    
    # Optional. If not specified then the name should reference a public stemcell.
    url = "https://bosh.io/d/stemcells/bosh-openstack-kvm-ubuntu-trusty-go_agent?v=2978"
    
    # Optional. If not specified then this will be determined by downloading the binary
    sha1 = "42e08d492dafd46226676be59ee3e6e8a0da618b" 
}
```

### "bosh_release"

Describes a bosh release that it can be referenced by bosh deployments.

```
resource "bosh_release" "bosh" {

    name = "bosh"
    url = "https://bosh.io/d/github.com/cloudfoundry/bosh?v=169"
    
    # Optional. If not specified then this will be determined by downloading the binary
    sha1 = "ec361150584094951273f1088344e2d4b2ebeb9f"
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

TODO

## Running Tests

You need to have a valid IaaS endpoint as well as a local bosh-lite instance to run the acceptance tests.

To run execute the following shell commands from the bosh provider directory.

```
$ export TF_ACC=1
$ eport BOSH_TARGET=#.#.#.#
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
