package cmd

import (
	"nagios-cli/nagios"

	"github.com/spf13/cobra"
)

var deleteHostCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a single or multiples Nagios host(s)",
	Long:  "Delete a single or multiples Nagios host(s)",
	Run: func(cmd *cobra.Command, args []string) {
		hosts := parseHosts()

		for _, host := range hosts {
			nagios.DeleteHost(nagiosConfig, host)
		}

		if applyConfig {
			nagios.ApplyConfig(nagiosConfig)
		}
	},
}

func init() {
	hostCmd.AddCommand(deleteHostCmd)
}
