package circleci

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
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
			"circleci_project": 				resourceCircleciProject(),
			"circleci_environment_variable": 	resourceCircleciEnvVar(),
		},
	}
	p.ConfigureFunc = providerConfiguretest(p)
	return p
}

func providerConfiguretest(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		config := Config{
			Token:        d.Get("token").(string),
			Organization: d.Get("organization").(string),
		}

		return config.Client()
	}
}
