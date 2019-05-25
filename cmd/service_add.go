package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var addServiceCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a single or multiples NagiosXI service(s)",
	Long:  "Add a single or multiples NagiosXI service(s)",
	Run: func(cmd *cobra.Command, args []string) {
		services := parseServices()

		for _, service := range services {
			nagiosxi.AddService(nagiosxiConfig, service)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	serviceCmd.AddCommand(addServiceCmd)
}
