package main

import (
	convox "github.com/mattaitchison/terraform-provider-convox/convox"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: convox.Provider,
	})
}
