package main

import (
    "github.com/hashicorp/terraform/plugin"
	"github.com/mevansam/terraform-provider-bosh/bosh"    
)

func main() {
	
	plugin.Serve( &plugin.ServeOpts {
		ProviderFunc: bosh.Provider,
	} )
}
