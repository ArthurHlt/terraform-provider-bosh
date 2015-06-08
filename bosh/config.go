package bosh

import (
	"golang.org/x/net/context"
	
	"github.com/mevansam/terraform-provider-bosh/bosh/bosh_client"
)

type Config struct {
	Target string
	User string
	Password string
}

func (c *Config) Client() (*bosh_client.Director, error) {
	
	b, err := bosh_client.NewDirector(context.Background(), c.Target, c.User, c.Password)
	if err != nil && err.Error() != "bosh director not found" {
		return nil, err
	}
	return b, nil
}
