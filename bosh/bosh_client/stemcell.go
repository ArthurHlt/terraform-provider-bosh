package bosh_client

import (
	"fmt"
)

type Stemcell struct {
	Name string
	Version string
	CID string
	Deployments []string
	
	api *Director
}

func (s *Stemcell) Delete() error {
	
	resp := s.api.director.DeleteStemcell(s.Name, s.Version)
	if resp.IsNotSuccessful() {
		return fmt.Errorf("error deleting stemcell '%s' with version '%s'", s.Name, s.Version)
	}
	return nil
}
