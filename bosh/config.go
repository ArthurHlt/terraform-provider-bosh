package bosh

import (
	"golang.org/x/net/context"
)

type Config struct {
	Target string
	User string
	Password string
	
	Verbose bool
}

func (c *Config) Client() (*BoshClient, error) {
	
	b, err := NewBoshClient(context.Background(), c.Target, c.User, c.Password, c.Verbose)
	if err != nil && err.Error() != "bosh director not found" {
		return nil, err
	}
	return b, nil
}
