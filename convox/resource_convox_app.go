package convox

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/convox/rack/client"
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceConvoxApp() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rack": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"environment": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"params": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},
			"formation": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"balancer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cpu": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"memory": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeInt},
						},
					},
				},
			},
		},
		Create: resourceConvoxAppCreate,
		Read:   resourceConvoxAppRead,
		Update: resourceConvoxAppUpdate,
		Delete: resourceConvoxAppDelete,
	}
}

func resourceConvoxAppCreate(d *schema.ResourceData, meta interface{}) error {
	c := RackClient(d, meta)
	if c == nil {
		return fmt.Errorf("Error rack client is nil: %#v", meta)
	}

	name := d.Get("name").(string)
	app, err := c.CreateApp(name)
	if err != nil {
		return fmt.Errorf(
			"Error creating app (%s): %s", name, err)
	}

	d.SetId(app.Name)
	stateConf := &resource.StateChangeConf{
		Pending: []string{"creating"},
		Target:  []string{"running"},
		Refresh: appRefreshFunc(c, app.Name),
		Timeout: 10 * time.Minute,
		Delay:   25 * time.Second,
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf(
			"Error waiting for app (%s) to be created: %s", app.Name, err)
	}
	return resourceConvoxAppUpdate(d, meta)
}

func resourceConvoxAppRead(d *schema.ResourceData, meta interface{}) error {
	c := RackClient(d, meta)
	app, err := c.GetApp(d.Get("name").(string))
	if err != nil {
		return err
	}
	d.SetId(app.Name)
	d.Set("release", app.Release)
	d.Set("status", app.Status)

	params, err := c.ListParameters(app.Name)
	if err != nil {
		return err
	}
	d.Set("params", params)

	env, err := c.GetEnvironment(app.Name)
	if err != nil {
		return err
	}
	d.Set("environment", env)

	formation, err := c.ListFormation(app.Name)
	if err != nil {
		return errwrap.Wrapf("Error while reading formation from Convox API: {{err}}", err)
	}
	return readFormation(d, formation)
}

func resourceConvoxAppUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)
	c := RackClient(d, meta)
	if err := setParams(c, d); err != nil {
		return err
	}
	d.SetPartial("params")

	if err := setEnv(c, d); err != nil {
		return err
	}
	d.SetPartial("environment")

	d.Partial(false)
	return resourceConvoxAppRead(d, meta)
}

func resourceConvoxAppDelete(d *schema.ResourceData, meta interface{}) error {
	c := RackClient(d, meta)
	_, err := c.DeleteApp(d.Id())
	return err
}

func readFormation(d *schema.ResourceData, v client.Formation) error {
	//formation := map[string]map[string]interface{}{}
	//// endpoints := []map[string]interface{}{}
	//for _, f := range v {
	//entry := map[string]interface{}{
	//"name":     f.Name,
	//"balancer": f.Balancer,
	//"cpu":      f.CPU,
	//"count":    f.Count,
	//"memory":   f.Memory,
	//"ports":    f.Ports,
	//}
	//formation[f.Name] = entry
	//// for _, port := range f.Ports {
	//// 	endpoints = append(endpoints, fmt.Sprintf("%s:%d (%s)", f.Balancer, port, f.Name))
	//// }
	//}

	//if err := d.Set("formation", formation); err != nil {
	//return errwrap.Wrapf("Unable to store formation: {{err}}", err)
	//}

	return nil
}

func setParams(c *client.Client, d *schema.ResourceData) error {
	if !d.HasChange("params") {
		return nil
	}

	raw := d.Get("params").(map[string]interface{})
	params := make(client.Parameters)
	for key := range raw {
		params[key] = raw[key].(string)
	}

	log.Printf("[DEBUG] Setting params: (%#v) for %s", params, d.Id())
	if err := c.SetParameters(d.Id(), params); err != nil {
		return fmt.Errorf("Error setting params (%#v) for %s: %s", params, d.Id(), err)
	}

	return nil
}

func setEnv(c *client.Client, d *schema.ResourceData) error {
	if !d.HasChange("environment") {
		return nil
	}

	env := d.Get("environment").(map[string]interface{})
	log.Printf("[DEBUG] Setting environment to (%#v) for %s", env, d.Id())
	data := ""
	for key, value := range env {
		data += fmt.Sprintf("%s=%s\n", key, value)
	}
	_, rel, err := c.SetEnvironment(d.Id(), strings.NewReader(data))
	if err != nil {
		return fmt.Errorf("Error setting vars (%#v) for %s: %s", env, d.Id(), err)
	}
	log.Printf("[INFO] Release (%s) created on: %s", rel, d.Id())

	return nil
}

func appRefreshFunc(c *client.Client, app string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		app, err := c.GetApp(app)
		if err != nil {
			return app, "", err
		}
		return app, app.Status, err
	}
}
