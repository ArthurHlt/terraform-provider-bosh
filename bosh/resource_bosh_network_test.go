package bosh

import (
//  "fmt"
    "testing"

    "github.com/hashicorp/terraform/helper/resource"
//  "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
)

func TestAccBoshNetwork_normal(t *testing.T) {

    resource.Test( t, 
        resource.TestCase {
            PreCheck: func() { testAccPreCheck(t) },
            Providers: testAccProviders,
            CheckDestroy: testAccCheckBoshNetworkDestroy,
            Steps: []resource.TestStep {
                },
        } )
}

func testAccCheckBoshNetworkDestroy(s *terraform.State) error {
    return nil
}
