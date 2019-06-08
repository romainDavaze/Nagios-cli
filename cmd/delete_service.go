package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteServicesCmd = &cobra.Command{
	Use:   "services <file>",
	Short: "Delete NagiosXI services",
	Long:  "Delete NagiosXI services",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		services := nagiosxi.ParseServices(objectsFile)

		for _, service := range services {
			nagiosxi.DeleteService(nagiosxiConfig, service)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteServicesCmd)
}
