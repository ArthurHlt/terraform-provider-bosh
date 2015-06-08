package bosh_client

import (
	"log"
	"testing"

	"github.com/kr/pretty"
)

func TestUploadReleases(t *testing.T) {

	var (
		err error
		release *Release
		releases map[string]*Release
	)
	
	d := GetDirector(t)
	
	releaselURLs := [...]string{
		"https://bosh.io/d/github.com/cf-platform-eng/docker-boshrelease?v=13",
		"https://bosh.io/d/github.com/cf-platform-eng/docker-boshrelease?v=12",
		"https://bosh.io/d/github.com/cloudfoundry-community/docker-registry-boshrelease?v=1",
	}
	
	for _, url := range releaselURLs {
		log.Printf("[DEBUG] Uploading release %s...", url)
		err = d.UploadRemoteRelease(url)
		if err != nil {
			log.Printf("[FAIL] Error uploading release: %s", err.Error())
			t.FailNow()
		}
	}
	
	releases, err = d.ListReleases()
	if len(releases) != len(releaselURLs) {
		log.Printf("[FAIL] Expected %d releases to have been uploaded but found %s.", len(releases), len(releaselURLs))
		t.FailNow()		
	}
	log.Printf("[DEBUG] Releases uploaded to Bosh: %# v", pretty.Formatter(releases))
	
	release = releases["docker/13"]
	if release == nil {
		log.Printf("[FAIL] Uploaded release 'docker-boshrelease' version '13' not returned from Bosh")
		t.FailNow()
	}
}

func TestDeleteReleases(t *testing.T) {

	d := GetDirector(t)
	releases, err := d.ListReleases()
	if err != nil {
		log.Printf("[FAIL] Error retrieving releases: %s", err.Error())
		t.FailNow()
	}

	for _, release := range releases {
		log.Printf("[DEBUG] Deleting release named '%s' and version '%s'...", release.Name, release.Version)
		err = release.Delete()
		if err != nil {
			log.Printf("[FAIL] Error deleting release: %s", err.Error())
			t.FailNow()
		}
	}
}
