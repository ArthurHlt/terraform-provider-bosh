package bosh_client

import (
	"fmt"
)

type Release struct {
	Name string
	Version string
	CommitHash string
	
	Deployed bool
	Uncommitted bool
	
	api *Director
}

func (r *Release) Delete() error {
	
	resp := r.api.director.DeleteRelease(r.Name, r.Version)
	if resp.IsNotSuccessful() {
		return fmt.Errorf("error deleting release '%s' with version '%s'", r.Name, r.Version)
	}
	return nil
}
