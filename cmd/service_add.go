package cmd

import (
	"nagios-cli/nagios"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a single or multiples Nagios service(s)",
	Long:  "Add a single or multiples Nagios service(s)",
	Run: func(cmd *cobra.Command, args []string) {
		services := parseServices()

		for _, service := range services {
			nagios.AddService(nagiosHost, apiKey, service)
		}
	},
}

func init() {
	serviceCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&nagiosFile, "file", "f", "", "file containing services to add")
	cobra.MarkFlagRequired(addCmd.Flags(), "file")
}
