package convox

import (
	"fmt"
	"log"

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
				Optional: true,
				Default:  false,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: ResourceConvoxSyslogCreateFactory(clientUnpacker),
		Read:   ResourceConvoxSyslogReadFactory(clientUnpacker),
		Update: ResourceConvoxSyslogUpdateFactory(clientUnpacker),
		Delete: ResourceConvoxSyslogDeleteFactory(clientUnpacker),
	}
}

func formURLString(d *schema.ResourceData) string {
	return fmt.Sprintf("%s://%s:%d", d.Get("scheme"), d.Get("hostname"), d.Get("port"))
}

// ResourceConvoxSyslogCreateFactory builds the Convox Syslog CreateFunc
func ResourceConvoxSyslogCreateFactory(clientUnpacker ClientUnpacker) schema.CreateFunc {
	log.Printf("ResourceConvoxSyslogCreateFactory")

	if clientUnpacker == nil {
		panic("clientUnpacker is required")
	}

	return func(d *schema.ResourceData, meta interface{}) error {
		log.Printf("CreateFunc")

		c, err := clientUnpacker(d, meta)
		if err != nil {
			return fmt.Errorf("Error unpacking client in CreateFunc: %s", err.Error())
		}

		options := map[string]string{
			"name": d.Get("name").(string),
			"Url":  formURLString(d),
		}

		if v, ok := d.GetOk("private"); ok {
			options["Private"] = fmt.Sprintf("%v", v)
		}

		log.Printf("[INFO] Calling Convox CreateResource...")
		_, err = c.CreateResource("syslog", options)
		if err != nil {
			return fmt.Errorf("Error calling CreateResource: %s -- %v", err.Error(), options)
		}

		// TODO: probably need to wait here for the status to stabilize. (and in update)

		d.Set("url", options["Url"])

		return nil
	}
}

// ResourceConvoxSyslogReadFactory builds the ReadFunc for the Convox Syslog Resource
func ResourceConvoxSyslogReadFactory(clientUnpacker ClientUnpacker) schema.ReadFunc {
	if clientUnpacker == nil {
		panic("clientUnpacker is required")
	}

	return func(d *schema.ResourceData, meta interface{}) error {
		c, err := clientUnpacker(d, meta)
		if err != nil {
			return fmt.Errorf("Error getting client in ReadFunc: %s", err.Error())
		}

		name := d.Get("name").(string)
		log.Printf("[INFO] Calling Convox GetResource(%s)...", name)
		convoxResource, err := c.GetResource(name)
		if err != nil {
			return fmt.Errorf("Error calling GetResource: %s", err.Error())
		}

		d.Set("url", convoxResource.Exports["URL"])

		return nil
	}
}

// ResourceConvoxSyslogUpdateFactory creates the UpdateFunc for the Convox Syslog Resource
func ResourceConvoxSyslogUpdateFactory(clientUnpacker ClientUnpacker) schema.UpdateFunc {
	if clientUnpacker == nil {
		panic("clientUnpacker is required")
	}

	return func(d *schema.ResourceData, meta interface{}) error {
		c, err := clientUnpacker(d, meta)
		if err != nil {
			return fmt.Errorf("Error getting client in UpdateFunc: %s", err.Error())
		}

		options := map[string]string{
			"Url": formURLString(d),
		}

		if v, ok := d.GetOk("private"); ok {
			options["Private"] = fmt.Sprintf("%v", v)
		}

		name := d.Get("name").(string)
		log.Printf("[INFO] Calling Convox UpdateResource(%s, <options>)...", name)
		_, err = c.UpdateResource(name, options)
		if err != nil {
			return fmt.Errorf("Error calling UpdateResource: %s -- %v", err.Error(), options)
		}

		d.Set("url", options["Url"])

		return nil
	}
}

// ResourceConvoxSyslogDeleteFactory builds the DeleteFunc for thw Convox Syslog Resource
func ResourceConvoxSyslogDeleteFactory(clientUnpacker ClientUnpacker) schema.DeleteFunc {
	if clientUnpacker == nil {
		panic("clientUnpacker is required")
	}

	return func(d *schema.ResourceData, meta interface{}) error {
		c, err := clientUnpacker(d, meta)
		if err != nil {
			return fmt.Errorf("Error getting client in DeleteFunc: %s", err.Error())
		}

		name := d.Get("name").(string)
		log.Printf("[INFO] Calling Convox DeleteResource(%s)...", name)
		_, err = c.DeleteResource(name)
		if err != nil {
			return fmt.Errorf("Error calling DeleteResource: %s", err.Error())
		}

		return nil
	}
}
