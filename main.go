package main

import (
	"github.com/form3tech-oss/terraform-provider-form3/form3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: form3.Provider,
	})
}
