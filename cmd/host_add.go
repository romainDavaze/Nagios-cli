package cmd

import (
	"nagios-cli/nagios"

	"github.com/spf13/cobra"
)

var addHostCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a single or multiples Nagios host(s)",
	Long:  "Add a single or multiples Nagios host(s)",
	Run: func(cmd *cobra.Command, args []string) {
		hosts := parseHosts()

		for _, host := range hosts {
			nagios.AddHost(nagiosConfig, host)
		}
	},
}

func init() {
	hostCmd.AddCommand(addHostCmd)

	addHostCmd.Flags().StringVarP(&nagiosFile, "file", "f", "", "file containing hosts to add")
	cobra.MarkFlagRequired(addHostCmd.Flags(), "file")
}
