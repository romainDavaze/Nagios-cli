package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var addHostCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a single or multiples NagiosXI host(s)",
	Long:  "Add a single or multiples NagiosXI host(s)",
	Run: func(cmd *cobra.Command, args []string) {
		hosts := parseHosts()

		for _, host := range hosts {
			nagiosxi.AddHost(nagiosxiConfig, host)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	hostCmd.AddCommand(addHostCmd)
}
