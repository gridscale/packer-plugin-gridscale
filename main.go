package main

import (
	"fmt"
	"os"

	"github.com/gridscale/packer-plugin-gridscale/builder/gridscale"
	"github.com/gridscale/packer-plugin-gridscale/version"
	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder(plugin.DEFAULT_NAME, new(gridscale.Builder))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
