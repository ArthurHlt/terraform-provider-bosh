package bosh

import (
//  "fmt"
    
//  "golang.org/x/net/context"
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceBoshJob() *schema.Resource {
    
    return &schema.Resource{
        
        Create: resourceBoshJobCreate,
        Read:   resourceBoshJobRead,
        Update: resourceBoshJobUpdate,
        Delete: resourceBoshJobDelete,

        Schema: map[string]*schema.Schema{
            
        },
    }
}

func resourceBoshJobCreate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshJobRead(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshJobUpdate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshJobDelete(d *schema.ResourceData, meta interface{}) error {
    return nil
}