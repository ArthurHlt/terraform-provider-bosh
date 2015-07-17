# Configure the OpenStack Provider
provider "openstack" {
    user_name  = "bosh"
    password = "bosh"
    tenant_name = "bosh_test_1"
    auth_url  = "https://os.pcf-services.com:5000/v2.0"
    insecure = true
    
    api_key = ""
    endpoint_type = ""
}

resource "openstack_compute_floatingip_v2" "bosh_target" {
    region = "durham"
    pool = "public01"
}
