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
		},
	}
}

func resourceCircleciProjectCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	organization := meta.(*Organization).name
	name := d.Get("name").(string)

	_, err := client.FollowProject(organization, name)
	if err != nil {
		return err
	}

	d.SetId(name)
	return nil
}

func resourceCircleciProjectRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCircleciProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCircleciProjectDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	organization := meta.(*Organization).name
	name := d.Get("name").(string)

	return client.DisableProject(organization, name)
}
