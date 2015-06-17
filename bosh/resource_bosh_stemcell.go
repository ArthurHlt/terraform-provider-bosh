package bosh

import (
	"fmt"
	"log"
    
    "github.com/hashicorp/terraform/helper/schema"

	"github.com/mevansam/terraform-provider-bosh/bosh/bosh_client"
)

func resourceBoshStemcell() *schema.Resource {
    
    return &schema.Resource{
        
        Create: resourceBoshStemcellCreate,
        Read:   resourceBoshStemcellRead,
        Delete: resourceBoshStemcellDelete,

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
			"cid": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
			},
		},
    }
}

func resourceBoshStemcellCreate(d *schema.ResourceData, meta interface{}) error {
	
	director := meta.(*bosh_client.Director)
	if director == nil {
		return fmt.Errorf("director is nil")
	}
	if !director.IsConnected() {
		return fmt.Errorf("director is not connected")
	}
	
	var (
		err error
		stemcell *bosh_client.Stemcell
	)
	
	url := d.Get("url").(string)
	log.Printf("[DEBUG] Uploading stemcell from '%s'.", url)
	
	err = director.UploadRemoteStemcell(url)
	if err != nil {
		return err
	}
	
	name := d.Get("name").(string)
	version := d.Get("version").(string)
	
	stemcell, err = director.GetStemcell(name, version)
	if err != nil {
		return fmt.Errorf("stemcell upload succeeded but unable to retrieve its details: %s", err.Error()) 
	}
	
	d.SetId(fmt.Sprintf("%s/%s", name, version))
	d.Set("cid", stemcell.CID)
    return nil
}

func resourceBoshStemcellRead(d *schema.ResourceData, meta interface{}) error {
	
	director := meta.(*bosh_client.Director)
	if director == nil {
		return fmt.Errorf("director is nil")
	}
	if !director.IsConnected() {
		return fmt.Errorf("director is not connected")
	}

	stemcell, err := director.GetStemcell(d.Get("name").(string), d.Get("version").(string))
	if err != nil {
		return err 
	}
	
	d.Set("cid", stemcell.CID)
    return nil
}

func resourceBoshStemcellDelete(d *schema.ResourceData, meta interface{}) error {
	
	director := meta.(*bosh_client.Director)
	if director == nil {
		return fmt.Errorf("director is nil")
	}
	if !director.IsConnected() {
		return fmt.Errorf("director is not connected")
	}
	
	stemcell, err := director.GetStemcell(d.Get("name").(string), d.Get("version").(string))
	if err != nil {
		return err 
	}

	err = stemcell.Delete()
	if err != nil {
		return err
	}
	
    return nil
}
