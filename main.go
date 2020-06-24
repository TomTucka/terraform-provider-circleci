package main

import (
	"github.com/TomTucka/terraform-provider-circleci/circleci"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return circleci.Provider()
		},
	})
}
