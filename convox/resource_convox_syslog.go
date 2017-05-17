package convox

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

// ResourceConvoxSyslog describes the schema for a Convox Syslog Resource resource
func ResourceConvoxSyslog(clientUnpacker ClientUnpacker) *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"scheme": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"private": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
				Default:  false,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: ResourceConvoxSyslogCreateFactory(clientUnpacker),
		Read:   ResourceConvoxSyslogReadFactory(clientUnpacker),
		//Update: ResourceConvoxSyslogUpdateFactory(clientUnpacker),
		// Delete: resourceConvoxSyslogDelete,
	}
}

// ResourceConvoxSyslogCreateFactory builds the Convox Syslog CreateFunc
func ResourceConvoxSyslogCreateFactory(clientUnpacker ClientUnpacker) schema.CreateFunc {
	if clientUnpacker == nil {
		panic("clientUnpacker is required")
	}

	return func(d *schema.ResourceData, meta interface{}) error {
		c, err := clientUnpacker(d, meta)
		if err != nil {
			return err
		}

		options := map[string]string{
			"url":     fmt.Sprintf("%s://%s:%d", d.Get("scheme"), d.Get("hostname"), d.Get("port")),
			"private": fmt.Sprintf("%v", d.Get("private")),
		}

		d.Set("url", options["url"])

		_, err = c.CreateResource("syslog", options)
		if err != nil {
			return err
		}

		return nil
	}
}

// ResourceConvoxSyslogReadFactory builds the ReadFunc for the Convox Syslog Resource
func ResourceConvoxSyslogReadFactory(clientUnpacker ClientUnpacker) schema.ReadFunc {
	if clientUnpacker == nil {
		panic("clientUnpacker is required")
	}

	return func(d *schema.ResourceData, meta interface{}) error {
		if d == nil {
			panic("d is required")
		}

		if meta == nil {
			panic("meta is required")
		}

		c, err := clientUnpacker(d, meta)
		if err != nil {
			return err
		}

		name := d.Get("name").(string)
		convoxResource, err := c.GetResource(name)
		if err != nil {
			return err
		}

		d.Set("url", convoxResource.Exports["URL"])

		return nil
	}
}
