package bosh

import (
//	"fmt"
	"log"
//	"os"
	"testing"
	
	"golang.org/x/net/context"
	"github.com/kr/pretty"
)

const boshLiteTarget = "192.168.50.4"
const boshLiteUser = "admin"
const boshLitePassword = "admin"

func TestBoshConnect(t *testing.T) {
	
	b, err := NewBoshClient(context.Background(), boshLiteTarget, boshLiteUser, boshLitePassword, testing.Verbose())
	if err != nil {
		log.Printf("[ERROR] Creating Bosh client failed: %s", err.Error())
		t.FailNow()
	}
	
	log.Printf("[DEBUG] Bosh client: %# v", pretty.Formatter(b))

	if !b.IsConnected() {
		log.Printf("[ERROR] Connection to Bosh lite target '%s' failed.", boshLiteTarget)
		t.FailNow()		
	}
	
	b, err = NewBoshClient(context.Background(), "127.0.0.1", boshLiteUser, boshLitePassword, testing.Verbose())
	if err == nil {
		log.Printf("[FAIL] Did not get expected bosh connection failure.")
		t.FailNow()
	} else if err.Error() != "bosh director not found" {
		log.Printf("[FAIL] Did not get expected bosh connection failure. Instead got error: %s", err.Error())
		t.FailNow()		
	}
	log.Printf("[DEBUG] Bosh connection failed as expected: %s", err.Error())
}
