package main

import (
	"log"

	// convox "github.com/GetTerminus/terraform-provider-convox/convox"

	"github.com/hashicorp/terraform/plugin"
	convox "github.com/terraform-provider-convox/convox"
)

var version = "0.1.5-dev"

func main() {
	log.Println("[INFO] convox provider version:", version)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: convox.Provider,
	})
}
