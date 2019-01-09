package convox

import (
	"io"

	"github.com/convox/rack/pkg/structs"
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
	AppCreate(name string, opts structs.AppCreateOptions) (*structs.App, error)
	AppGet(name string) (*structs.App, error)
	AppDelete(name string) error

	ListFormation(app string) (structs.Formation, error)

	ListParameters(app string) (map[string]string, error)
	SetParameters(app string, params map[string]string) error

	GetEnvironment(app string) (structs.Environment, error)
	SetEnvironment(app string, body io.Reader) (structs.Environment, string, error)

	ResourceCreate(kind string, opts structs.ResourceCreateOptions) (*structs.Resource, error)
	ResourceDelete(name string) error
	ResourceGet(name string) (*structs.Resource, error)
	ResourceUpdate(name string, opts structs.ResourceUpdateOptions) (*structs.Resource, error)

	ResourceLink(name, app string) (*structs.Resource, error)
	ResourceUnlink(name, app string) (*structs.Resource, error)
}
