package bosh_client

import (
	"fmt"
	"strings"
	
	"golang.org/x/net/context"
	"github.com/cloudfoundry-community/gogobosh"
	"github.com/cloudfoundry-community/gogobosh/api"
	"github.com/cloudfoundry-community/gogobosh/net"
)

type Director struct {
	UUID string
	Name string
	Version string
	CPI string
	
	director api.DirectorRepository
}

func NewDirector(ctx context.Context, target string, user string, password string) (*Director, error) {
	
	d := gogobosh.NewDirector(target, user, password)	
	c := &Director {
		director: api.NewBoshDirectorRepository(&d, net.NewDirectorGateway()),
	}
	
	err := c.Connect()
	if err != nil && !strings.Contains(err.Error(), "connection refused") {
		return nil, err
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
