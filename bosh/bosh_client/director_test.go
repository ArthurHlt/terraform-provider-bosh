package bosh_client

import (
	"log"
	"regexp"
	"strings"
	"testing"
	
	"golang.org/x/net/context"
	"github.com/kr/pretty"
)

const BoshLiteTarget = "https://127.0.0.1:25555"
const BoshLiteUser = "admin"
const BoshLitePassword = "admin"

func TestDirectorConnect(t *testing.T) {

	d := GetDirector(t)
	log.Printf("[DEBUG] Bosh client: %# v", pretty.Formatter(d))

	if !d.IsConnected() {
		log.Printf("[FAIL] Connection to Bosh lite target '%s' failed.", BoshLiteTarget)
		t.FailNow()		
	}
}

func TestDirectorConnectFailure(t *testing.T) {

	d, err := NewDirector(context.Background(), "https://127.0.0.1:25554", BoshLiteUser, BoshLitePassword)
	if d.IsConnected() {
		log.Printf("[FAIL] Expected initial connection to non-existent Bosh target 'https://127.0.0.1:25554' to fail.")
		t.FailNow()		
	}
	
	err = d.Connect()
	if err == nil {
		log.Printf("[FAIL] Expected bosh connection failure when trying to connect explictly to 'https://127.0.0.1:25554'")
		t.FailNow()
	} else if !strings.Contains(err.Error(), "connection refused") {
		log.Printf("[FAIL] Did not get expected bosh connection failure. Instead got error: %s", err.Error())
		t.FailNow()		
	}
	log.Printf("[DEBUG] Bosh connection failed as expected with error '%s'", err.Error())
}

func GetDirector(t *testing.T) *Director {
	
	d, err := NewDirector(context.Background(), BoshLiteTarget, BoshLiteUser, BoshLitePassword)
	if err != nil {
		log.Printf("[FAIL] Unable to connnect to Bosh director: %s", err.Error())
		t.FailNow()
		return nil
	}
	
	return d	
}

func GetInitializedDirector(t *testing.T) *Director {

	var (
		err error
		
		stemcell *Stemcell
		release *Release
	)
	
	d := GetDirector(t)
	if d != nil {
		
		stemcell, err = d.GetStemcell("bosh-warden-boshlite-ubuntu-trusty-go_agent", "2776")
		if err != nil {
			log.Printf("[FAIL] Unable to retrieve lookup stemcells: %s", err.Error())
			t.FailNow()
			return nil
		}
		if stemcell == nil {
			log.Printf("[DEBUG] Uploading stemcell for testing...")
			err = d.UploadRemoteStemcell("https://bosh.io/d/stemcells/bosh-warden-boshlite-ubuntu-trusty-go_agent?v=2776")
			if err != nil {
				log.Printf("[FAIL] Error uploading stemcell: %s", err.Error())
				t.FailNow()
				return nil
			}			
		}
		
		release, err = d.GetRelease("docker", "13")
		if err != nil {
			log.Printf("[FAIL] Unable to retrieve lookup releases: %s", err.Error())
			t.FailNow()			
			return nil
		}
		if release == nil {
			log.Printf("[DEBUG] Uploading release for testing...")
			err = d.UploadRemoteRelease("https://bosh.io/d/github.com/cf-platform-eng/docker-boshrelease?v=13")
			if err != nil {
				log.Printf("[FAIL] Error uploading release: %s", err.Error())				
			}
		}
	}
	return d
}

func assertPattern(t *testing.T, s string, p string) {
	
	r := regexp.MustCompile(p)
	if !r.MatchString(s) {
		log.Printf("[FAIL] Pattern '%s' not found in string not as expected.", p)
		t.FailNow()		
	}
}

func assertPatternFalse(t *testing.T, s string, p string) {
	
	r := regexp.MustCompile(p)
	if r.MatchString(s) {
		log.Printf("[FAIL] Pattern '%s' found in string not as expected.", p)
		t.FailNow()		
	}
}
