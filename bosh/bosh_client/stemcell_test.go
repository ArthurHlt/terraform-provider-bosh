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
		"https://bosh.io/d/stemcells/bosh-openstack-kvm-ubuntu-trusty-go_agent-raw?v=2978",
//		"https://bosh.io/d/stemcells/bosh-openstack-kvm-ubuntu-trusty-go_agent-raw?v=2977",
//		"https://bosh.io/d/stemcells/bosh-warden-boshlite-ubuntu-trusty-go_agent?v=2776",
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
	if err != nil {
		log.Printf("[FAIL] Error retrieving stemcells: %s", err.Error())
		t.FailNow()
	}
	if len(stemcells) <= len(stemcellURLs) {
		log.Printf("[FAIL] Expected %d stemcells to have been uploaded but found %d.", len(stemcellURLs), len(stemcells))
		t.FailNow()		
	}
	log.Printf("[DEBUG] Stemcells uploaded to Bosh: %# v", pretty.Formatter(stemcells))
	
	stemcell = stemcells["bosh-openstack-kvm-ubuntu-trusty-go_agent-raw/2978"]
	if stemcell == nil {
		log.Printf("[FAIL] Uploaded stemcell 'bosh-openstack-kvm-ubuntu-trusty-go_agent-raw' version '2978' not returned from Bosh")
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

	stemcells, err = d.ListStemcells()
	if len(stemcells) != 0 {
		log.Printf("[FAIL] Expected all stemcells to have been deleted but found %d stemcells remaining.", len(stemcells))
		t.FailNow()		
	}
}
