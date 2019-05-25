package cmd

import (
	"github.com/romainDavaze/nagios-cli/nagios"
	"github.com/spf13/cobra"
)

var deleteServiceCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a single or multiples Nagios service(s)",
	Long:  "Delete a single or multiples Nagios service(s)",
	Run: func(cmd *cobra.Command, args []string) {
		services := parseServices()

		for _, service := range services {
			nagios.DeleteService(nagiosConfig, service)
		}

		if applyConfig {
			nagios.ApplyConfig(nagiosConfig)
		}
	},
}

func init() {
	serviceCmd.AddCommand(deleteServiceCmd)
}
