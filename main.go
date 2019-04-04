package main

import (
	"fmt"

	"github.com/HewlettPackard/HPE-OneSphere-Terraform-Provider/onesphereterraform"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: onesphereterraform.Provider})
	fmt.Println("In main.go file")

}
