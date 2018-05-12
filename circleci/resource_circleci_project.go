package circleci

import (
	"github.com/hashicorp/terraform/helper/schema"
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
			},
		},
	}
}

func resourceCircleciProjectCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	organization := meta.(*Organization).name
	project := d.Get("name").(string)
	env_vars := d.Get("env_vars").(map[string]interface{})

	_, err := client.FollowProject(organization, project)
	if err != nil {
		return err
	}

	d.SetId(project)

	for name, value := range env_vars {
		_, err := client.AddEnvVar(organization, project, name, value.(string))
		if err != nil {
			return err
		}
	}
	return nil
}

func resourceCircleciProjectRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCircleciProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceCircleciProjectCreate(d, meta)
}

func resourceCircleciProjectDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	organization := meta.(*Organization).name
	name := d.Get("name").(string)

	return client.DisableProject(organization, name)
}
