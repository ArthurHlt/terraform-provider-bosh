package bosh

import (
//  "fmt"
    
//  "golang.org/x/net/context"
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceBoshDeployment() *schema.Resource {
    
    return &schema.Resource{
        
        Create: resourceBoshDeploymentCreate,
        Read:   resourceBoshDeploymentRead,
        Update: resourceBoshDeploymentUpdate,
        Delete: resourceBoshDeploymentDelete,

        Schema: map[string]*schema.Schema{
            
        },
    }
}

func resourceBoshDeploymentCreate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshDeploymentRead(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshDeploymentUpdate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceBoshDeploymentDelete(d *schema.ResourceData, meta interface{}) error {
    return nil
}