package bosh

import (
	"golang.org/x/net/context"
)

type Config struct {
	Target string
	User string
	Password string
}

func (c *Config) Client() (*BoshClient, error) {
	
	b, err := NewBoshClient(context.Background(), c.Target, c.User, c.Password)
	if err != nil && err.Error() != "bosh director not found" {
		return nil, err
	}
	return b, nil
}
