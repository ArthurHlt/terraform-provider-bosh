package bosh

import (
//  "fmt"
    "testing"

    "github.com/hashicorp/terraform/helper/resource"
//  "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
)

func TestAccBoshManifestTemplate_normal(t *testing.T) {

    resource.Test( t, 
        resource.TestCase {
            PreCheck: func() { testAccPreCheck(t) },
            Providers: testAccProviders,
            CheckDestroy: testAccCheckBoshManifestTemplateDestroy,
            Steps: []resource.TestStep {
                },
        } )
}

func testAccCheckBoshManifestTemplateDestroy(s *terraform.State) error {
    return nil
}
