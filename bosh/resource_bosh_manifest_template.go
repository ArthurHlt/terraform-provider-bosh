package bosh

import (
//  "fmt"
    
//  "golang.org/x/net/context"
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceBoshManifestTemplate() *schema.Resource {
    
    return &schema.Resource{
        
        Create: resourceBoshManifestTemplateCreate,
        Read:   resourceBoshManifestTemplateRead,
        Update: resourceBoshManifestTemplateUpdate,
        Delete: resourceBoshManifestTemplateDelete,

        Schema: map[string]*schema.Schema{
            
        },
    }
}

func resourceBoshManifestTemplateCreate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshManifestTemplateRead(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshManifestTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshManifestTemplateDelete(d *schema.ResourceData, meta interface{}) error {
    return nil
}