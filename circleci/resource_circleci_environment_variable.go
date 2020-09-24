package circleci

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCircleciEnvVar() *schema.Resource {
	return &schema.Resource{
		Create: resourceCircleciEnvVarCreate,
		Read:   resourceCircleciEnvVarRead,
		Update: resourceCircleciEnvVarUpdate,
		Delete: resourceCircleciEnvVarDelete,
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
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceCircleciEnvVarCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	organization := meta.(*Organization).name
	vcs := meta.(*Organization).VCS
	project := d.Get("project").(string)
	name := d.Get("name").(string)
	value := d.Get("value").(string)
	_, err := client.AddEnvVar(vcs, organization, project, name, value)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s-%s", project, name))
	return nil
}

func resourceCircleciEnvVarRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCircleciEnvVarUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceCircleciEnvVarCreate(d, meta)
}

func resourceCircleciEnvVarDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	organization := meta.(*Organization).name
	vcs := meta.(*Organization).VCS
	project := d.Get("project").(string)
	name := d.Get("name").(string)
	return client.DeleteEnvVar(vcs, organization, project, name)
}
