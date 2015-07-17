package bosh

import (
//  "fmt"
    
//  "golang.org/x/net/context"
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceBoshMicrobosh() *schema.Resource {
    
    return &schema.Resource{
        
        Create: resourceBoshMicroboshCreate,
        Read:   resourceBoshMicroboshRead,
        Update: resourceBoshMicroboshUpdate,
        Delete: resourceBoshMicroboshDelete,

        Schema: map[string]*schema.Schema{
            
        },
    }
}

func resourceBoshMicroboshCreate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshMicroboshRead(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshMicroboshUpdate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshMicroboshDelete(d *schema.ResourceData, meta interface{}) error {
    return nil
}