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

func (d *Director) UploadRemoteRelease(url string) error {
	
	resp := d.director.UploadRemoteRelease(url)
	if resp.IsNotSuccessful() {
		return fmt.Errorf("error uploading release at %s: %s", url, resp.Message)
	}
	return nil
}

func (d *Director) ListReleases() (map[string]*Release, error) {
	
	releases := make(map[string]*Release)

	releaseList, resp := d.director.GetReleases()
	if resp.IsNotSuccessful() {
		return nil, fmt.Errorf("Could not fetch BOSH releases: %s", resp.Message)
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

func (d *Director) GetRelease(name string, version string) (*Release, error) {
	
	releases, err := d.ListReleases()
	if err != nil {
		return nil, err
	}
	
	return releases[name + "/" + version], nil	
}

func (r *Release) Delete() error {
	
	resp := r.api.director.DeleteRelease(r.Name, r.Version)
	if resp.IsNotSuccessful() {
		return fmt.Errorf("error deleting release '%s' with version '%s'", r.Name, r.Version)
	}
	return nil
}
