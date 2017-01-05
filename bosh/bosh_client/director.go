package bosh_client

import (
	"log"
	"fmt"
	"regexp"
	"strings"
	
	"golang.org/x/net/context"
	"github.com/cloudfoundry-community/gogobosh"
)

type Director struct {
	UUID string
	Name string
	Version string
	CPI string
	
	director gogobosh.DirectorRepository
}

func NewDirector(ctx context.Context, target string, user string, password string) (*Director, error) {

	var re *regexp.Regexp

	re = regexp.MustCompile("^http(s)?://")
	if re.FindString(target) == "" {
		target = "https://" + target
	}
	re = regexp.MustCompile(":\\d+$")
	if re.FindString(target) == "" {
		target = target + ":25555"
	}
	
	d := gogobosh.NewDirector(target, user, password)	
	c := &Director {
		director: api.NewBoshDirectorRepository(&d, net.NewDirectorGateway()),
	}
	
	err := c.Connect()
	if err != nil {
		if !strings.Contains(err.Error(), "connection refused") {
			return nil, err
		}
		log.Printf("[DEBUG] Connection to director not successful. Could be because " + 
			"the Director does no exist yet. So connection has been deferred.")
	} else {
		log.Printf("[DEBUG] Connection to director successful.")
	}
	
	return c, nil 
}

func (d *Director) Connect() error {
	
	info, resp := d.director.GetInfo()
	if resp.IsNotSuccessful() {
		return fmt.Errorf("directory information query was unsuccessful: %s - %s", resp.ErrorCode, resp.Message)
	}
	
	d.UUID = info.UUID
	d.Name = info.Name
	d.Version = info.Version
	d.CPI = info.CPI
	
	return nil
}

func (b *Director) IsConnected() bool {
	return b.UUID != ""
}
