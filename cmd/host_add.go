package cmd

import (
	"github.com/romainDavaze/nagios-cli/nagios"
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

		if applyConfig {
			nagios.ApplyConfig(nagiosConfig)
		}
	},
}

func init() {
	hostCmd.AddCommand(addHostCmd)
}
