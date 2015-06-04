package cloudfoundry

import (
//	"fmt"
//	"net/url"
)

type Config struct {
	ApiEndpoint string
	Username string
	Password string
}

func (c *Config) Client() (*interface{}, error) {
	return nil, nil
}
