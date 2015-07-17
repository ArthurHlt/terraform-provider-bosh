# Configure the OpenStack Provider
provider "bosh" {
    target  = "127.0.0.1"
    user  = "admin"
    password = "admin"
}

resource "bosh_stemcell" "ubuntu-warden" {
	
	name = "bosh-warden-boshlite-ubuntu-trusty-go_agent"
	version = "2776"
	url = "https://bosh.io/d/stemcells/bosh-warden-boshlite-ubuntu-trusty-go_agent?v=2776"
}

#resource "bosh_release" "docker" {
#	
#	name = "docker"
#	version = "13"
#	url = "https://bosh.io/d/github.com/cf-platform-eng/docker-boshrelease?v=13"
#}
