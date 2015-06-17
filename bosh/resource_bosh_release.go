package bosh

import (
	"fmt"
	"log"
    
    "github.com/hashicorp/terraform/helper/schema"

	"github.com/mevansam/terraform-provider-bosh/bosh/bosh_client"
)

func resourceBoshRelease() *schema.Resource {
    
    return &schema.Resource{
        
        Create: resourceBoshReleaseCreate,
        Read:   resourceBoshReleaseRead,
        Delete: resourceBoshReleaseDelete,

		Schema: map[string]*schema.Schema{
			
			"name": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"version": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"url": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},			
			"sha1": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},			
			"commit_hash": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
			},
		},
    }
}

func resourceBoshReleaseCreate(d *schema.ResourceData, meta interface{}) error {
	
	director := meta.(*bosh_client.Director)
	if director == nil {
		return fmt.Errorf("director is nil")
	}
	if !director.IsConnected() {
		return fmt.Errorf("director is not connected")
	}

	var (
		err error
		release *bosh_client.Release
	)
	
	url := d.Get("url").(string)
	log.Printf("[DEBUG] Uploading release from '%s'.", url)
	
	err = director.UploadRemoteRelease(url)
	if err != nil {
		return err
	}
	
	name := d.Get("name").(string)
	version := d.Get("version").(string)
	
	release, err = director.GetRelease(name, version)
	if err != nil {
		return fmt.Errorf("release upload succeeded but unable to retrieve its details: %s", err.Error()) 
	}
	
	d.SetId(fmt.Sprintf("%s/%s", name, version))
	d.Set("commit_hash", release.CommitHash)
    return nil
}

func resourceBoshReleaseRead(d *schema.ResourceData, meta interface{}) error {
	
	director := meta.(*bosh_client.Director)
	if director == nil {
		return fmt.Errorf("director is nil")
	}
	if !director.IsConnected() {
		return fmt.Errorf("director is not connected")
	}

	release, err := director.GetRelease(d.Get("name").(string), d.Get("version").(string))
	if err != nil {
		return err 
	}
	
	d.Set("commit_hash", release.CommitHash)
    return nil
}

func resourceBoshReleaseDelete(d *schema.ResourceData, meta interface{}) error {
	
	director := meta.(*bosh_client.Director)
	if director == nil {
		return fmt.Errorf("director is nil")
	}
	if !director.IsConnected() {
		return fmt.Errorf("director is not connected")
	}
	
	release, err := director.GetRelease(d.Get("name").(string), d.Get("version").(string))
	if err != nil {
		return err 
	}

	err = release.Delete()
	if err != nil {
		return err
	}
	
    return nil
}
