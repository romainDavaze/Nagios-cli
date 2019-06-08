package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var addServicesCmd = &cobra.Command{
	Use:   "services <file>",
	Short: "Add NagiosXI services",
	Long:  "Add NagiosXI services",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		services := nagiosxi.ParseServices(objectsFile)

		for _, service := range services {
			nagiosxi.AddService(nagiosxiConfig, service, force)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	addCmd.AddCommand(addServicesCmd)
}
