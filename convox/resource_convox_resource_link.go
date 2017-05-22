package convox

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceConvoxResourceLink(clientUnpacker ClientUnpacker) *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rack": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Create: ResourceConvoxResourceLinkCreateFactory(clientUnpacker),
		Delete: ResourceConvoxResourceLinkDeleteFactory(clientUnpacker),
	}
}

func ResourceConvoxResourceLinkCreateFactory(clientUnpacker ClientUnpacker) schema.CreateFunc {

	return func(d *schema.ResourceData, meta interface{}) error {
		if clientUnpacker == nil {
			return errors.New("clientUnpacker is required")
		}

		c, err := clientUnpacker(d, meta)
		if err != nil {
			return fmt.Errorf("Error unpacking client in CreateFunc: %s", err.Error())
		}

		resource := d.Get("resource_name").(string)
		app := d.Get("app_name").(string)

		_, err = c.CreateLink(resource, app)
		if err != nil {
			return fmt.Errorf("Error calling CreateLink(%s, %s): %s", resource, app, err.Error())
		}

		d.SetId(fmt.Sprintf("%s-%s", resource, app))

		return nil
	}
}

func ResourceConvoxResourceLinkDeleteFactory(clientUnpacker ClientUnpacker) schema.DeleteFunc {
	return func(d *schema.ResourceData, meta interface{}) error {
		if clientUnpacker == nil {
			return errors.New("clientUnpacker is required")
		}

		c, err := clientUnpacker(d, meta)
		if err != nil {
			return fmt.Errorf("Error unpacking client in DeleteFunc: %s", err.Error())
		}

		resource := d.Get("resource_name").(string)
		app := d.Get("app_name").(string)

		_, err = c.DeleteLink(resource, app)
		if err != nil {
			return fmt.Errorf("Error calling DeleteLink(%s, %s): %s", resource, app, err.Error())
		}

		return nil
	}
}
