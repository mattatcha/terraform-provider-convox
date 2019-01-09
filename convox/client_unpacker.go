package convox

import "github.com/convox/rack/pkg/structs"

// ClientUnpacker unpacks the client provided in the meta-data object
type ClientUnpacker func(valueGetter ValueGetter, meta interface{}) (structs.Provider, error)
