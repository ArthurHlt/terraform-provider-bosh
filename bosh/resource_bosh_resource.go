package bosh

import (
//  "fmt"
    
//  "golang.org/x/net/context"
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceBoshResource() *schema.Resource {
    
    return &schema.Resource{
        
        Create: resourceBoshResourceCreate,
        Read:   resourceBoshResourceRead,
        Update: resourceBoshResourceUpdate,
        Delete: resourceBoshResourceDelete,

        Schema: map[string]*schema.Schema{
            
        },
    }
}

func resourceBoshResourceCreate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshResourceRead(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshResourceUpdate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshResourceDelete(d *schema.ResourceData, meta interface{}) error {
    return nil
}