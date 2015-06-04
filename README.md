# Terraform Bosh Provider

The [Bosh](http://bosh.io/) provider for [Terraform](https://terraform.io/) provides a seamless integration of the Bosh deployment and operations toolset with a Terraform'ed cloud.

## Provider

The Bosh provider will target the given endpoint if it is available. If not the "bosh-deploy" resource can be used to build the micro-bosh deployment to be targeted.

```
provider "bosh" {
	target = "#.#.#.#"	# Publicly addressable name or IP
    username = "admin"
    password = "*****"
}
```

## Terraform Resources

### "bosh_microbosh"

Deploys Microbosh to the specified IaaS using *[bosh-init](https://github.com/cloudfoundry/bosh-init)*. If this resource is specified and the ```target``` given in the provider configuration is not pingable, then this resource will attempt to create the Microbosh instance.

```
resource "bosh_microbosh" {

	name = "my_terraformed_microbosh"
	
    binary = true  # set up microbosh in a binary configuration for HA

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

```
GLOG_logtostderr=1
TF_ACC=1
```