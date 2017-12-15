package main

import (
	"log"

	"github.com/hashicorp/terraform/plugin"
	convox "github.com/mattatcha/terraform-provider-convox/convox"
)

var version = "0.1.7-dev"

func main() {
	log.Println("[INFO] convox provider version:", version)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: convox.Provider,
	})
}
