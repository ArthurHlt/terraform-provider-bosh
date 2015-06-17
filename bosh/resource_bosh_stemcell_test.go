package bosh

import (
	"log"
	"fmt"
    "testing"

    "github.com/hashicorp/terraform/helper/resource"
    "github.com/hashicorp/terraform/terraform"
	"github.com/kr/pretty"
	
	"github.com/mevansam/terraform-provider-bosh/bosh/bosh_client"
)

const testStemcellResourceRef = "bosh_stemcell.ubuntu"
const testStemcellName = "bosh-openstack-kvm-ubuntu-trusty-go_agent"
const testStemcellVersion = "2978"

const testAccBoshStemcellConfig = `

resource "bosh_stemcell" "ubuntu" {
	
	name = "` + testStemcellName + `"
	version = "` + testStemcellVersion + `"
	url = "https://bosh.io/d/stemcells/bosh-openstack-kvm-ubuntu-trusty-go_agent?v=2978"
	sha1 = "42e08d492dafd46226676be59ee3e6e8a0da618b" 
}
`

func TestAccBoshStemcell_normal(t *testing.T) {

    resource.Test( t, 
        resource.TestCase {
            PreCheck: func() { testAccPreCheck(t) },
            Providers: testAccProviders,
            CheckDestroy: testAccCheckBoshStemcellDestroy(testStemcellResourceRef, testStemcellName, testStemcellVersion),
            Steps: []resource.TestStep {
					resource.TestStep {
						Config: testAccBoshStemcellConfig,
						Check: resource.ComposeTestCheckFunc(
							
							testAccCheckBoshStemcellExists(testStemcellResourceRef),							
							resource.TestCheckResourceAttr(
								"bosh_stemcell.ubuntu", "name", testStemcellName),
							resource.TestCheckResourceAttr(
								"bosh_stemcell.ubuntu", "version", testStemcellVersion),
							resource.TestCheckResourceAttr(
								"bosh_stemcell.ubuntu", "url", "https://bosh.io/d/stemcells/bosh-openstack-kvm-ubuntu-trusty-go_agent?v=2978"),
							resource.TestCheckResourceAttr(
								"bosh_stemcell.ubuntu", "sha1", "42e08d492dafd46226676be59ee3e6e8a0da618b"),
						),
					},
                },
        } )
}

func testAccCheckBoshStemcellExists(resource string) resource.TestCheckFunc {
	
	return func(s *terraform.State) error {
		
		var (
			err error
			
			director *bosh_client.Director
			stemcell *bosh_client.Stemcell
		)
		
		director, err = GetDirector()
		if err != nil {
			return err
		}

		rs, ok := s.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("stemcell resource '%s' not found in terraform state", resource)
		}
		
		log.Printf("[DEBUG] Resource state: %# v", pretty.Formatter(rs))
		
		attributes := rs.Primary.Attributes		
		stemcell, err = director.GetStemcell(attributes["name"], attributes["version"])
		if err != nil {
			return err
		}
		if stemcell.CID != attributes["cid"] {
			return fmt.Errorf("retrieved stemcell CID '%s' does not match '%s' which is " + 
				"the CID of resource state '%s'", stemcell.CID, attributes["cid"], resource)
		}
		
		return nil
	}
}

func testAccCheckBoshStemcellDestroy(resource string, stemcellName string, stemcellVersion string) resource.TestCheckFunc {
	
	return  func(s *terraform.State) error {
		
		_, ok := s.RootModule().Resources[resource]
		if ok {
			return fmt.Errorf("stemcell resource '%s' still exists in the terraform state", resource)
		}
		
		var (
			err error
			
			director *bosh_client.Director
			stemcell *bosh_client.Stemcell
		)
		
		director, err = GetDirector()
		if err != nil {
			return err
		}
		
		stemcell, err = director.GetStemcell(stemcellName, stemcellVersion)
		if err != nil {
			return err
		}
		if stemcell != nil {
			return fmt.Errorf("stemcell '%s' version '%s' was not deleted as expected", stemcellName, stemcellVersion)
		}
		
		return nil
	}
}
