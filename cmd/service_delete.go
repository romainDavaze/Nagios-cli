package cmd

import (
	"nagios-cli/nagios"

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
	},
}

func init() {
	serviceCmd.AddCommand(deleteServiceCmd)

	deleteServiceCmd.Flags().StringVarP(&nagiosFile, "file", "f", "", "file containing services to delete")
	cobra.MarkFlagRequired(deleteServiceCmd.Flags(), "file")
}
