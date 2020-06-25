package circleci

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCircleciProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceCircleciProjectCreate,
		Read:   resourceCircleciProjectRead,
		Update: resourceCircleciProjectUpdate,
		Delete: resourceCircleciProjectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"env_vars": {
				Type:     schema.TypeMap,
				Optional: true,
				Sensitive: true
			},
		},
	}
}

func resourceCircleciProjectCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Organization).client
	organization := m.(*Organization).name
	project := d.Get("name").(string)
	envVars := d.Get("env_vars").(map[string]interface{})

	_, err := client.FollowProject(organization, project)
	if err != nil {
		return err
	}

	d.SetId(project)

	for name, value := range envVars {
		_, err := client.AddEnvVar(organization, project, name, value.(string))
		if err != nil {
			return err
		}
	}
	return nil
}

func resourceCircleciProjectRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCircleciProjectUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCircleciProjectCreate(d, m)
}

func resourceCircleciProjectDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Organization).client
	organization := m.(*Organization).name
	name := d.Get("name").(string)

	return client.DisableProject(organization, name)
}
