package convox

import (
	"sync"

	"github.com/convox/rack/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

const clientVersion = "20160910125708"

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"rack": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
					},
				},
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
		// "convox_rack": dataSourceConvoxRack(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"convox_app": resourceConvoxApp(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	racks := d.Get("rack").(*schema.Set).List()
	client := &Client{racks: make(map[string]*client.Client)}
	for _, rack := range racks {
		m := rack.(map[string]interface{})
		client.Set(m["name"].(string), m["hostname"].(string), m["password"].(string))
	}
	return client, nil
}

type Client struct {
	mu    sync.Mutex
	racks map[string]*client.Client
}

func (c *Client) Set(name, host, password string) *client.Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.racks[name] = client.New(host, password, clientVersion)

	return c.racks[name]
}

func (c *Client) Rack(name string) *client.Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.racks[name]
}

func rackClient(d *schema.ResourceData, meta interface{}) *client.Client {
	return meta.(*Client).Rack(d.Get("rack").(string))
}
