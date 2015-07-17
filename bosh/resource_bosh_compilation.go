package bosh

import (
//  "fmt"
    
//  "golang.org/x/net/context"
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceBoshCompilation() *schema.Resource {
    
    return &schema.Resource{
        
        Create: resourceBoshCompilationCreate,
        Read:   resourceBoshCompilationRead,
        Update: resourceBoshCompilationUpdate,
        Delete: resourceBoshCompilationDelete,

        Schema: map[string]*schema.Schema{
            
        },
    }
}

func resourceBoshCompilationCreate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshCompilationRead(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshCompilationUpdate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshCompilationDelete(d *schema.ResourceData, meta interface{}) error {
    return nil
}