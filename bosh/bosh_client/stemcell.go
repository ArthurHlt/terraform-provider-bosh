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

func (d *Director) UploadRemoteStemcell(url string) error {
	
	resp := d.director.UploadRemoteStemcell(url)
	if resp.IsNotSuccessful() {
		return fmt.Errorf("error uploading stemcell at %s: %s", url, resp.Message)
	}
	return nil
}

func (d *Director) ListStemcells() (map[string]*Stemcell, error) {
	
	stemcells := make(map[string]*Stemcell)

	stemcellList, resp := d.director.GetStemcells()
	if resp.IsNotSuccessful() {
		return nil, fmt.Errorf("Could not fetch BOSH stemcells: %s", resp.Message)
	} else {
		for _, s := range stemcellList {
			stemcell := Stemcell{
				Name: s.Name,
				Version: s.Version,
				CID: s.Cid,
				Deployments: s.Deployments,
				api: d,
			}
			stemcells[s.Name + "/" + s.Version] = &stemcell
		}
	}
	
	return stemcells, nil
}

func (d *Director) GetStemcell(name string, version string) (*Stemcell, error) {
	
	stemcells, err := d.ListStemcells()
	if err != nil {
		return nil, err
	}
	
	return stemcells[name + "/" + version], nil	
}

func (s *Stemcell) Delete() error {
	
	resp := s.api.director.DeleteStemcell(s.Name, s.Version)
	if resp.IsNotSuccessful() {
		return fmt.Errorf("error deleting stemcell '%s' with version '%s'", s.Name, s.Version)
	}
	return nil
}
