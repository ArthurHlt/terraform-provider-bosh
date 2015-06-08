package bosh_client

import (
	"fmt"
	"strings"
	
	"golang.org/x/net/context"
	"github.com/cloudfoundry-community/gogobosh"
	"github.com/cloudfoundry-community/gogobosh/api"
//	"github.com/cloudfoundry-community/gogobosh/models"
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

func (d *Director) UploadRemoteStemcell(url string) error {
	
	resp := d.director.UploadRemoteStemcell(url)
	if resp.IsNotSuccessful() {
		return fmt.Errorf("error uploading stemcell at %s", url)
	}
	return nil
}

func (d *Director) ListStemcells() (map[string]*Stemcell, error) {
	
	stemcells := make(map[string]*Stemcell)

	stemcellList, resp := d.director.GetStemcells()
	if resp.IsNotSuccessful() {
		return nil, fmt.Errorf("Could not fetch BOSH stemcells")
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


func (d *Director) UploadRemoteRelease(url string) error {
	
	resp := d.director.UploadRemoteRelease(url)
	if resp.IsNotSuccessful() {
		return fmt.Errorf("error uploading stemcell at %s", url)
	}
	return nil
}

func (d *Director) ListReleases() (map[string]*Release, error) {
	
	releases := make(map[string]*Release)

	releaseList, resp := d.director.GetReleases()
	if resp.IsNotSuccessful() {
		return nil, fmt.Errorf("Could not fetch BOSH stemcells")
	} else {
		for _, r := range releaseList {
			for _, v := range r.Versions {
				release := Release{
					Name: r.Name,
					Version: v.Version,
					CommitHash: v.CommitHash,
					Deployed: v.CurrentlyDeployed,
					Uncommitted: v.UncommittedChanges,
					api: d,
				}
				
				releases[r.Name + "/" + v.Version] = &release
			}
		}
	}
	
	return releases, nil
}
