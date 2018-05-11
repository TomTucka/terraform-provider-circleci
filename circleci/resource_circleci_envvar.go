package circleci

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCircleciEnvvar() *schema.Resource {
	return &schema.Resource{
		Create: resourceCircleciEnvvarCreate,
		Read:   resourceCircleciEnvvarRead,
		Update: resourceCircleciEnvvarUpdate,
		Delete: resourceCircleciEnvvarDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCircleciEnvvarCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	organization := meta.(*Organization).name
	project := d.Get("project").(string)
	name := d.Get("name").(string)
	value := d.Get("value").(string)

	_, err := client.AddEnvVar(organization, project, name, value)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s-%s", project, name))
	return nil
}

func resourceCircleciEnvvarRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCircleciEnvvarUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceCircleciEnvvarCreate(d, meta)
}

func resourceCircleciEnvvarDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	organization := meta.(*Organization).name
	project := d.Get("project").(string)
	name := d.Get("name").(string)

	return client.DeleteEnvVar(organization, project, name)
}
