package circleci

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CIRCLECI_TOKEN", nil),
				Description: "CircleCI API token.",
			},
			"organization": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CIRCLECI_ORGANIZATION", nil),
				Description: "CircleCI organization name to manage.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"example_server": resourceCircleciProject(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return Client(Config{
		Token:        d.Get("api_token").(string),
		Organization: d.Get("organization").(string),
	})
}