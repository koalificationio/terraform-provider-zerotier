package main

import (
	"github.com/koalificationio/terraform-provider-zerotier/zerotier"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: zerotier.Provider})
}
