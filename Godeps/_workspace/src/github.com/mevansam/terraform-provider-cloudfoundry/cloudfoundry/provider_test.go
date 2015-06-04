package cloudfoundry

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"vsphere": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	apiEndpoint := os.Getenv("CF_API_ENDPOINT")
	username := os.Getenv("CF_USERNAME")
	password := os.Getenv("CF_PASSWORD")
	if apiEndpoint == "" || username == "" || password == "" {
		t.Fatal("CF_API_ENDPOINT, CF_USERNAME and CF_PASSWORD must be set for acceptance tests to work.")
	}
}
