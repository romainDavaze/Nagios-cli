package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteServiceCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a single or multiples NagiosXI service(s)",
	Long:  "Delete a single or multiples NagiosXI service(s)",
	Run: func(cmd *cobra.Command, args []string) {
		services := parseServices()

		for _, service := range services {
			nagiosxi.DeleteService(nagiosxiConfig, service)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	serviceCmd.AddCommand(deleteServiceCmd)
}
