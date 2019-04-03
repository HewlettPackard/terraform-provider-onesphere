package main

import (
	"fmt"

	"./onesphereterraform"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: onesphereterraform.Provider})
	fmt.Println("In main.go file")

}
