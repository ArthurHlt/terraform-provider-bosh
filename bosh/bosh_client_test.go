package bosh

import (
	"log"
	"strings"
	"testing"
	
	"golang.org/x/net/context"
	"github.com/kr/pretty"
)

const boshLiteTarget = "https://192.168.50.4:25555"
const boshLiteUser = "admin"
const boshLitePassword = "admin"

func TestBoshConnect(t *testing.T) {
	
	b, err := NewBoshClient(context.Background(), boshLiteTarget, boshLiteUser, boshLitePassword)
	if err != nil {
		log.Printf("[FAIL] Creating Bosh client failed: %s", err.Error())
		t.FailNow()
	}
	
	log.Printf("[DEBUG] Bosh client: %# v", pretty.Formatter(b))

	if !b.IsConnected() {
		log.Printf("[FAIL] Connection to Bosh lite target '%s' failed.", boshLiteTarget)
		t.FailNow()		
	}
	
	b, err = NewBoshClient(context.Background(), "https://127.0.0.1:25555", boshLiteUser, boshLitePassword)
	if b.IsConnected() {
		log.Printf("[FAIL] Expected initial connection to non-existent Bosh target 'https://127.0.0.1:25555' to fail.")
		t.FailNow()		
	}
	
	err = b.Connect()
	if err == nil {
		log.Printf("[FAIL] Expected bosh connection failure when trying to connect explictly to 'https://127.0.0.1:25555'")
		t.FailNow()
	} else if !strings.Contains(err.Error(), "connection refused") {
		log.Printf("[FAIL] Did not get expected bosh connection failure. Instead got error: %s", err.Error())
		t.FailNow()		
	}
	log.Printf("[DEBUG] Bosh connection failed as expected: %s", err.Error())
}
