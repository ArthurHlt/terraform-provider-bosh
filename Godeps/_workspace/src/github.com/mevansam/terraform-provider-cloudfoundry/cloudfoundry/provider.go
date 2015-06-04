package cloudfoundry

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"cf_api_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CF_API_ENDPOINT", nil),
			},
			"cf_username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CF_USERNAME", nil),
			},
			"cf_password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CF_PASSWORD", nil),
			},
		},

		// TODO:

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		ApiEndpoint: d.Get("cf_api_endpoint").(string),
		Username: d.Get("vsphere_password").(string),
		Password: d.Get("vsphere_host").(string),
	}
	return config.Client()
}
