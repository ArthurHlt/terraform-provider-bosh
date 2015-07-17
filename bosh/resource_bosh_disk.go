package bosh

import (
//  "fmt"
    
//  "golang.org/x/net/context"
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceBoshDisk() *schema.Resource {
    
    return &schema.Resource{
        
        Create: resourceBoshDiskCreate,
        Read:   resourceBoshDiskRead,
        Update: resourceBoshDiskUpdate,
        Delete: resourceBoshDiskDelete,

        Schema: map[string]*schema.Schema{
            
        },
    }
}

func resourceBoshDiskCreate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshDiskRead(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshDiskUpdate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshDiskDelete(d *schema.ResourceData, meta interface{}) error {
    return nil
}