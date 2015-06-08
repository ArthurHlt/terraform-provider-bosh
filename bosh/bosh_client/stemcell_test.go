package bosh_client

import (
	"log"
	"testing"

	"github.com/kr/pretty"
)

func TestUploadStemcells(t *testing.T) {

	var (
		err error
		stemcell *Stemcell
		stemcells map[string]*Stemcell
	)
	
	d := GetDirector(t)
	
	stemcellURLs := [...]string{
		"https://bosh.io/d/stemcells/bosh-warden-boshlite-ubuntu-trusty-go_agent?v=2776",
//		"https://bosh.io/d/stemcells/bosh-openstack-kvm-ubuntu-trusty-go_agent-raw?v=2978",
//		"https://bosh.io/d/stemcells/bosh-openstack-kvm-ubuntu-trusty-go_agent-raw?v=2977",
	}
	
	for _, url := range stemcellURLs {
		log.Printf("[DEBUG] Uploading stemcell %s...", url)
		err = d.UploadRemoteStemcell(url)
		if err != nil {
			log.Printf("[FAIL] Error uploading stemcell: %s", err.Error())
			t.FailNow()
		}
	}
	
	stemcells, err = d.ListStemcells()
	if len(stemcells) != len(stemcellURLs) {
		log.Printf("[FAIL] Expected %d stemcells to have been uploaded but found %s.", len(stemcellURLs), len(stemcells))
		t.FailNow()		
	}
	log.Printf("[DEBUG] Stemcells uploaded to Bosh: %# v", pretty.Formatter(stemcells))
	
	stemcell = stemcells["bosh-warden-boshlite-ubuntu-trusty-go_agent/2776"]
	if stemcell == nil {
		log.Printf("[FAIL] Uploaded stemcell 'bosh-warden-boshlite-ubuntu-trusty-go_agent' version '2776' not returned from Bosh")
		t.FailNow()
	}
}

func TestDeleteStemcells(t *testing.T) {
	
	d := GetDirector(t)
	stemcells, err := d.ListStemcells()
	if err != nil {
		log.Printf("[FAIL] Error retrieving stemcells: %s", err.Error())
		t.FailNow()
	}

	for _, stemcell := range stemcells {
		log.Printf("[DEBUG] Deleting stemcell named '%s' and version '%s'...", stemcell.Name, stemcell.Version)
		err = stemcell.Delete()
		if err != nil {
			log.Printf("[FAIL] Error deleting stemcell: %s", err.Error())
			t.FailNow()
		}
	}
}
