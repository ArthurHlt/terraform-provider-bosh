package bosh_client

import (
	"log"
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
	}
	
	return d	
}
