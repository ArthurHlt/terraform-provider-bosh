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

const testReleaseResourceRef = "bosh_release.docker"
const testReleaseName = "docker"
const testReleaseVersion = "12"

const testAccBoshReleaseConfig = `

resource "bosh_release" "docker" {
	
	name = "` + testReleaseName + `"
	version = "` + testReleaseVersion + `"
	url = "https://bosh.io/d/github.com/cf-platform-eng/docker-boshrelease?v=12"
	sha1 = "5e382b11379246bb351174c61e09992fc350019f" 
}
`

func TestAccBoshRelease_normal(t *testing.T) {

    resource.Test( t, 
        resource.TestCase {
            PreCheck: func() { testAccPreCheck(t) },
            Providers: testAccProviders,
            CheckDestroy: testAccCheckBoshReleaseDestroy(testReleaseResourceRef, testReleaseName, testReleaseVersion),
            Steps: []resource.TestStep {
					resource.TestStep {
						Config: testAccBoshReleaseConfig,
						Check: resource.ComposeTestCheckFunc(
							
							testAccCheckBoshReleaseExists(testReleaseResourceRef),							
							resource.TestCheckResourceAttr(
								"bosh_release.docker", "name", testReleaseName),
							resource.TestCheckResourceAttr(
								"bosh_release.docker", "version", testReleaseVersion),
							resource.TestCheckResourceAttr(
								"bosh_release.docker", "url", "https://bosh.io/d/github.com/cf-platform-eng/docker-boshrelease?v=12"),
							resource.TestCheckResourceAttr(
								"bosh_release.docker", "sha1", "5e382b11379246bb351174c61e09992fc350019f"),
						),
					},
                },
        } )
}

func testAccCheckBoshReleaseExists(resource string) resource.TestCheckFunc {
	
	return func(s *terraform.State) error {
		
		var (
			err error
			
			director *bosh_client.Director
			release *bosh_client.Release
		)
		
		director, err = GetDirector()
		if err != nil {
			return err
		}

		rs, ok := s.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("release resource '%s' not found in terraform state", resource)
		}
		
		log.Printf("[DEBUG] Resource state: %# v", pretty.Formatter(rs))
		
		attributes := rs.Primary.Attributes		
		release, err = director.GetRelease(attributes["name"], attributes["version"])
		if err != nil {
			return err
		}
		if release.CommitHash != attributes["commit_hash"] {
			return fmt.Errorf("retrieved release commite hash '%s' does not match '%s' which is " + 
				"the commit hash of resource state '%s'", release.CommitHash, attributes["commit_hash"], resource)
		}
		
		return nil
	}
}

func testAccCheckBoshReleaseDestroy(resource string, releaseName string, releaseVersion string) resource.TestCheckFunc {
	
	return  func(s *terraform.State) error {
		
		_, ok := s.RootModule().Resources[resource]
		if ok {
			return fmt.Errorf("release resource '%s' still exists in the terraform state", resource)
		}
		
		var (
			err error
			
			director *bosh_client.Director
			release *bosh_client.Release
		)
		
		director, err = GetDirector()
		if err != nil {
			return err
		}
		
		release, err = director.GetRelease(releaseName, releaseVersion)
		if err != nil {
			return err
		}
		if release != nil {
			return fmt.Errorf("release '%s' version '%s' was not deleted as expected", releaseName, releaseVersion)
		}
		
		return nil
	}
}
