package bosh

import (	
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"target": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BOSH_TARGET", nil),
			},
			"user": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BOSH_USER", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BOSH_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"bosh_stemcell": resourceBoshStemcell(),
			"bosh_release": resourceBoshRelease(),
			"bosh_cloudconfig": resourceBoshCloudConfig(),
			"bosh_deployment": resourceBoshDeployment(),
			"bosh_microbosh": resourceBoshMicrobosh(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config := Config{
		Target: d.Get("target").(string)	,
		User: d.Get("user").(string),
		Password: d.Get("password").(string),
	}
	return config.Client()
}
