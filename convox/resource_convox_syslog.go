package convox

import (
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
				Type:     schema.TypeString,
				Required: true,
			},
			"scheme": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
		Create: ResourceConvoxSyslogCreateFactory(clientUnpacker),
		// Read:   resourceConvoxSyslogRead,
		// Update: resourceConvoxSyslogUpdate,
		// Delete: resourceConvoxSyslogDelete,
	}
}

// ResourceConvoxSyslogCreateFactory builds the Convox Syslog CreateFunc
func ResourceConvoxSyslogCreateFactory(clientUnpacker ClientUnpacker) schema.CreateFunc {
	if clientUnpacker == nil {
		panic("clientUnpacker is required")
	}

	return func(d *schema.ResourceData, meta interface{}) error {
		_, err := clientUnpacker(d, meta)
		if err != nil {
			return err
		}

		return nil
		// if c == nil {
		// 	return fmt.Errorf("Error rack client is nil: %#v", meta)
		// }

		// name := d.Get("name").(string)

		// c.CreateResource()

		// app, err := c.CreateApp(name)
		// if err != nil {
		// 	return fmt.Errorf(
		// 		"Error creating app (%s): %s", name, err)
		// }

		// d.SetId(app.Name)
		// stateConf := &resource.StateChangeConf{
		// 	Pending: []string{"creating"},
		// 	Target:  []string{"running"},
		// 	Refresh: appRefreshFunc(c, app.Name),
		// 	Timeout: 10 * time.Minute,
		// 	Delay:   25 * time.Second,
		// }

		// if _, err = stateConf.WaitForState(); err != nil {
		// 	return fmt.Errorf(
		// 		"Error waiting for app (%s) to be created: %s", app.Name, err)
		// }
		// return resourceConvoxAppUpdate(d, meta)
	}
}
