package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var addServicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Add NagiosXI services",
	Long:  "Add NagiosXI services",
	Run: func(cmd *cobra.Command, args []string) {
		services := nagiosxi.ParseServices(objectsFile)

		for _, service := range services {
			nagiosxi.AddService(nagiosxiConfig, service)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	addCmd.AddCommand(addServicesCmd)
}
