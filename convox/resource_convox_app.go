package convox

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/convox/rack/client"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/serenize/snaker"
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
			"params": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"deployment_maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
						},
						"deployment_minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
						},
						"environment": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"internal": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
						},
						"key": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"private": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
						},
						"release": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"repository": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"security_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"subnets": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnets_private": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vpc": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vpc_cidr": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		Read:   resourceConvoxAppRead,
		Create: resourceConvoxAppCreate,
		Update: resourceConvoxAppUpdate,
		Delete: resourceConvoxAppDelete,
	}
}

func resourceConvoxAppCreate(d *schema.ResourceData, meta interface{}) error {
	c := rackClient(d, meta)
	if c == nil {
		return fmt.Errorf("client nil: %+v", meta)
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
		Refresh: createAppRefreshFunc(c, app.Name),
		Timeout: 10 * time.Minute,
		Delay:   25 * time.Second,
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf(
			"Error waiting for app (%s) to be created: %s", app.Name, err)
	}
	return resourceConvoxAppUpdate(d, meta)
}

func resourceConvoxAppUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)
	c := rackClient(d, meta)
	if err := setParams(c, d); err != nil {
		return err
	}
	d.SetPartial("params")

	d.Partial(false)
	return resourceConvoxAppRead(d, meta)
}

func resourceConvoxAppDelete(d *schema.ResourceData, meta interface{}) error {
	c := rackClient(d, meta)
	_, err := c.DeleteApp(d.Id())
	return err
}
func resourceConvoxAppRead(d *schema.ResourceData, meta interface{}) error {
	c := rackClient(d, meta)
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
	return d.Set("params", []interface{}{paramsToMap(params)})
}

func setParams(c *client.Client, d *schema.ResourceData) error {
	if d.HasChange("params") {
		oraw, nraw := d.GetChange("params")

		// FIXME: oraw and nraw could be nil!
		o := oraw.([]interface{})[0].(map[string]interface{})
		n := nraw.([]interface{})[0].(map[string]interface{})
		params := paramsFromMap(diffParams(o, n))
		if len(params) > 0 {
			log.Printf("[DEBUG] Setting params: (%#v) for %s", params, d.Id())

			if err := c.SetParameters(d.Id(), params); err != nil {
				return fmt.Errorf("Error setting params (%#v) for %s: %s", params, d.Id(), err)
			}
		}
	}
	return nil
}
func paramsFromMap(m map[string]interface{}) client.Parameters {
	result := make(client.Parameters)
	for k, v := range m {
		k = snaker.SnakeToCamel(k)
		if k == "Vpc" || k == "VpcCidr" {
			k = strings.ToUpper(k)
		}
		// FIXME: won't work when params supports complex types.
		result[k] = fmt.Sprint(v)
	}
	return result
}

func paramsToMap(ps client.Parameters) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range ps {
		k = toSnake(k)
		switch k {
		case "internal", "private":
			if strings.ToLower(v) == "no" {
				result[k] = false
				continue
			}
			result[k] = true
		case "deployment_minimum", "deployment_maximum":
			i, _ := strconv.Atoi(v)
			result[k] = i
		case "vpccidr":
			result["vpc_cidr"] = v
		default:
			result[k] = v
		}
	}
	return result
}

func diffParams(oldParams, newParams map[string]interface{}) map[string]interface{} {
	changed := make(map[string]interface{})
	for k, v := range newParams {
		old, ok := oldParams[k]
		if ok && old != v {
			changed[k] = v
		}
	}
	return changed
}

func createAppRefreshFunc(c *client.Client, app string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		app, err := c.GetApp(app)
		if err != nil {
			return app, "", err
		}
		return app, app.Status, err
	}
}
