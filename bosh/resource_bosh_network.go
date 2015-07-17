package bosh

import (
//  "fmt"
    
//  "golang.org/x/net/context"
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceBoshNetwork() *schema.Resource {
    
    return &schema.Resource{
        
        Create: resourceBoshNetworkCreate,
        Read:   resourceBoshNetworkRead,
        Update: resourceBoshNetworkUpdate,
        Delete: resourceBoshNetworkDelete,

        Schema: map[string]*schema.Schema{
            
        },
    }
}

func resourceBoshNetworkCreate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshNetworkRead(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshNetworkDelete(d *schema.ResourceData, meta interface{}) error {
    return nil
}