# Configure the OpenStack Provider
provider "bosh" {
    target  = "${openstack_compute_floatingip_v2.bosh_target.address}"
    user  = "admin"
    password = "admin"
}

resource "bosh_stemcell" "ubuntu-kvm-raw" {
	
	name = "bosh-openstack-kvm-ubuntu-trusty-go_agent-raw"
	version = "2986"
	url = "https://bosh.io/d/stemcells/bosh-openstack-kvm-ubuntu-trusty-go_agent-raw?v=2986"
}

resource "bosh_release" "docker" {
	
	name = "docker"
	version = "13"
	url = "https://bosh.io/d/github.com/cf-platform-eng/docker-boshrelease?v=13"
}
