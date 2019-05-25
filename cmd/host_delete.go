package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteHostCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a single or multiples NagiosXI host(s)",
	Long:  "Delete a single or multiples NagiosXI host(s)",
	Run: func(cmd *cobra.Command, args []string) {
		hosts := parseHosts()

		for _, host := range hosts {
			nagiosxi.DeleteHost(nagiosxiConfig, host)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	hostCmd.AddCommand(deleteHostCmd)
}
