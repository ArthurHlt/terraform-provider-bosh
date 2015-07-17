package bosh

import (
//  "fmt"
    
//  "golang.org/x/net/context"
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceBoshCloudConfig() *schema.Resource {
    
    return &schema.Resource{
        
        Create: resourceBoshCloudConfigCreate,
        Read:   resourceBoshCloudConfigRead,
        Update: resourceBoshCloudConfigUpdate,

        Schema: map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
            
			"network": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
						"type": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
						"cidr": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
						},
						"gateway": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
						},
						"num_reserved_ips": &schema.Schema{
							Type: schema.TypeInt,
							Optional: true,
						},
						"static_ip_block": &schema.Schema{
							Type:     schema.TypeSet,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type: schema.TypeString,
										Required: true,
									},
									"num_ips": &schema.Schema{
										Type: schema.TypeInt,
										Required: true,
									},
								},
							},
						},
						"cloud_property": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type: schema.TypeString,
										Required: true,
									},
									"value": &schema.Schema{
										Type: schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
//				Set: resourceCloudConfigNetworkHash,
			},            
        },
    }
}

func resourceBoshCloudConfigCreate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshCloudConfigRead(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshCloudConfigUpdate(d *schema.ResourceData, meta interface{}) error {
    return nil
}
