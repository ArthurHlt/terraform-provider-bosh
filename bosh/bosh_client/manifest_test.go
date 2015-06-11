package bosh_client

import (
	"log"
	"testing"
	"github.com/kr/pretty"
)

func TestManifestProcessing(t *testing.T) {
	
	d := GetDirector(t)
	m, err := d.NewManifest(testTemplate)
	if err != nil {
		log.Printf("[FATAL] Error creating manifest instance: %s", err.Error())
		t.FailNow()
	}
	
	props := m.Properties
	
	props["key1a"] = "value1a"
	props["key1b"] = "value1b"
	
	props["nested1a"] = make(map[string]interface{})
	nested1a := props["nested1a"].(map[string]interface{})
	nested1a["key2a"] = "value2a"
	
	nested1a["nested2a"] = make(map[string]interface{})
	nested2a := nested1a["nested2a"].(map[string]interface{})
	nested2a["key3a"] = "value3a"
	
	nested1a["array2a"] = []interface{}{ "array2a_1", "array2a_2", "array2a_3" } 
	
	props["key1c"] = "value1c"
	
	props["nested1b"] = make(map[string]interface{})
	nested1b := props["nested1b"].(map[string]interface{})
	
	nested1b["array2b"] = []interface{}{ make(map[string]interface{}), make(map[string]interface{}) }
	array2b := nested1b["array2b"].([]interface{})
	array2b1 := array2b[0].(map[string]interface{})
	array2b2 := array2b[1].(map[string]interface{})
	
	array2b1["key2b11"] = "value2b11"
	array2b1["key2b12"] = "value2b12"
	array2b2["key2b21"] = "value2b21"
	
	props["key1d"] = "value1d"
	
	log.Printf("[DEBUG] Test properties: %# v", pretty.Formatter(props))

	s, err := m.Process()
	if err != nil {
		log.Printf("[FATAL] Error processing manifest: %s", err.Error())
		t.FailNow()
	}
	
	log.Printf("[DEBUG] Processed manifest:\n%s", s)
	
	assertPattern(t, s, "UUID: [0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12} ;")
	assertPattern(t, s, "UUID: [0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12} ;")
	assertPattern(t, s, "key1a: value1a ;")
	assertPattern(t, s, "key1b: value1b ;")
	assertPattern(t, s, "key1c: value1c ;")
	assertPattern(t, s, "key1d: value1d ;")
	assertPattern(t, s, "key1z:  ;")
	assertPattern(t, s, "key2a: value2a ;")
	assertPattern(t, s, "key2b:  ;")
	assertPattern(t, s, "key3a: value3a ;")
	assertPattern(t, s, "array2a\\[0\\]: array2a_1 ;")
	assertPattern(t, s, "array2a\\[1\\]: array2a_2 ;")
	assertPattern(t, s, "array2a\\[2\\]: array2a_3 ;")
	assertPattern(t, s, "key2b11: value2b11;")
	assertPattern(t, s, "key2b12: value2b12;")
	assertPattern(t, s, "key2b21: value2b21;")
}

const testTemplate = `
UUID: {{ .DirectorUUID }} ;
LatestStemcellVersion: {{ .LatestStemcellVersion "bosh-warden-boshlite-ubuntu-trusty-go_agent" }} ;
LatestReleaseVersion: {{ .LatestReleaseVersion "docker" }} ;
key1a: {{ .P "key1a" }} ;
key1b: {{ .P "key1b" }} ;
key1c: {{ .P "key1c" }} ;
key1d: {{ .P "key1d" }} ;
key1z: {{ .P "key1z" }} ;
key2a: {{ .P "nested1a.key2a" }} ;
key2b: {{ .P "nested1a.key2b" }} ;
key3a: {{ .P "nested1a.nested2a.key3a" }} ;

array2a: {{ range .P "nested1a.array2a" }}
	{{ . }}; 
{{ end }}
array2a[0]: {{ .P "nested1a.array2a[0]" }} ;
array2a[1]: {{ .P "nested1a.array2a[1]" }} ;
array2a[2]: {{ .P "nested1a.array2a[2]" }} ;

array2b: {{ range $index, $elmt := .P "nested1b.array2b" }}
	{{ $elmt }}; {{ if (eq $index 0) }}
		key2b11: {{ PP $elmt "key2b11" }};
		key2b12: {{ PP $elmt "key2b12" }}; 
	{{ end }} {{ if (eq $index 1) }}
		key2b21: {{ PP $elmt "key2b21" }};
	{{ end }} {{ end }}
key2b11: {{ .P "nested1b.array2b[0].key2b11" }};
key2b12: {{ .P "nested1b.array2b[0].key2b12" }}; 
key2b21: {{ .P "nested1b.array2b[1].key2b21" }};
`
