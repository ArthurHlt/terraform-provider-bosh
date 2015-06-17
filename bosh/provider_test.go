package bosh

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	
	"github.com/mevansam/terraform-provider-bosh/bosh/bosh_client"
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

func GetDirector() (*bosh_client.Director, error) {
	
	d, err := bosh_client.NewDirector(context.Background(), os.Getenv("BOSH_TARGET"), os.Getenv("BOSH_USER"), os.Getenv("BOSH_PASSWORD"))
	if err != nil {
		return nil, fmt.Errorf("error connecting to the bosh director: %s", err.Error())
	}
	
	return d, nil
}
