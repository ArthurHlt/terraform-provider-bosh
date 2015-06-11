package bosh_client

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strconv"
	"strings"
	"reflect"
	"regexp"
)

type Manifest struct {

	Properties map[string]interface{}

	releases map[string]*Release
	stemcells map[string]*Stemcell
		
	template string
	
	api *Director
}

func (d *Director) NewManifest(template string) (*Manifest, error) {

	var (
		err error
	)
	
	m := &Manifest {
		Properties: make(map[string]interface{}),
		template: template,
		api: d,
	}
	
	m.releases, err = d.ListReleases()
	if err != nil {
		return nil, err
	}
	m.stemcells, err = d.ListStemcells()
	if err != nil {
		return nil, err
	}
	
	return m, nil
}

func (m *Manifest) Process() (string, error) {
	
	var (
		err error
		out bytes.Buffer
	)
	
	t := template.New("Manifest Template")
	t = t.Funcs(template.FuncMap{"E": lookupEnvironment})
	t = t.Funcs(template.FuncMap{"PP": lookupProperty})
	
	t, err = t.Parse(m.template)
	if err != nil {
		return "", fmt.Errorf("[FAIL] Error parsing deployment manifest template: %s", err.Error())
	}
	
	err = t.Execute(&out, m)
	if err != nil {
		return "", fmt.Errorf("[FAIL] Error processing deployment manifest template: %s", err.Error())
	}
	
	return out.String(), nil
}

func (m *Manifest) DirectorUUID() string {
	return m.api.UUID
}

func (m *Manifest) LatestStemcellVersion(name string) string {
	
	var (
		version string
	)
	
	for _, s := range m.stemcells {
		if s.Name == name && s.Version > version {
			version = s.Version
		}
	}
	
	return version
}

func (m *Manifest) LatestReleaseVersion(name string) string {

	var (
		version string
	)
	
	for _, s := range m.releases {
		if s.Name == name && s.Version > version {
			version = s.Version
		}
	}

	return version
}

func (m *Manifest) P(key string) (interface{}, error) {
	
	keys := strings.Split(key, ".")
	return lookup(0, &keys, m.Properties)
}

func lookupEnvironment(variable string) string {
	return os.Getenv(variable)
}

func lookupProperty(properties map[string]interface{}, key string) (interface{}, error) {
	
	keys := strings.Split(key, ".")
	return lookup(0, &keys, properties)
}

func lookup(i int, keys *[]string, properties map[string]interface{}) (interface{}, error) {
	
	s := (*keys)[i]
	r := regexp.MustCompile("^([-+_0-9a-zA-Z]+)(\\[(\\d+)\\])?$")
	if !r.MatchString(s) {
		return nil, fmt.Errorf("Invalid key %s", s)		
	}
	k := r.FindStringSubmatch(s)
	
	key := k[1]
	index := -1
	if k[3] != "" {
		index, _ = strconv.Atoi(k[3])
	}
	
	v := properties[key]
	if v == nil {
		return nil, nil
	}
	
	t := reflect.TypeOf(v).Kind()
	if index != -1 {
		if t == reflect.Slice || t == reflect.Array {
			v = (v.([]interface{}))[index]
			t = reflect.TypeOf(v).Kind()
		} else {
			return nil, fmt.Errorf("Error key %s is not an array", key)
		}
	}
	if (i == len(*keys)-1) {
		return v, nil
	}
	if t == reflect.Map {
		return lookup(i+1, keys, v.(map[string]interface{}))
	}
	return nil, fmt.Errorf("Value at key %s is not a map", s)
}
