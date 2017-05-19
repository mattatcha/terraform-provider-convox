package convox

import (
	"io"

	"github.com/convox/rack/client"
)

// ValueGetter gets values (from terraform schema)
type ValueGetter interface {
	Get(key string) interface{}
	GetOk(key string) (interface{}, bool)
}

// ClientUnpacker unpacks the client provided in the meta-data object
type ClientUnpacker func(valueGetter ValueGetter, meta interface{}) (Client, error)

// Client interface is the subset of the convox client we use
type Client interface {
	CreateApp(name string) (*client.App, error)
	GetApp(name string) (*client.App, error)
	DeleteApp(name string) (*client.App, error)

	ListFormation(app string) (client.Formation, error)

	ListParameters(app string) (client.Parameters, error)
	SetParameters(app string, params map[string]string) error

	GetEnvironment(app string) (client.Environment, error)
	SetEnvironment(app string, body io.Reader) (client.Environment, string, error)

	GetResource(name string) (*client.Resource, error)
	CreateResource(kind string, options map[string]string) (*client.Resource, error)
	UpdateResource(name string, options map[string]string) (*client.Resource, error)
	DeleteResource(name string) (*client.Resource, error)
}
