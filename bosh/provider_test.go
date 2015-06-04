package bosh

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
		"bosh": testAccProvider,
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
	target := os.Getenv("BOSH_TARGET")
	user := os.Getenv("BOSH_USER")
	password := os.Getenv("BOSH_PASSWORD")
	if target == "" || user == "" || password == "" {
		t.Fatal("BOSH_TARGET, BOSH_USER and BOSH_PASSWORD must be set for acceptance tests to work.")
	}
}
