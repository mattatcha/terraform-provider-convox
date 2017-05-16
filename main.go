package main

import (
	"log"

	convox "github.com/mattaitchison/terraform-provider-convox/convox"

	"github.com/hashicorp/terraform/plugin"
)

var version = "0.1.5-dev"

func main() {
	log.Println("[INFO] convox provider version:", version)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: convox.Provider,
	})
}
