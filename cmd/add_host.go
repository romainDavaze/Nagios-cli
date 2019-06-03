package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var addHostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "Add NagiosXI hosts",
	Long:  "Add NagiosXI hosts",
	Run: func(cmd *cobra.Command, args []string) {
		hosts := nagiosxi.ParseHosts(objectsFile)

		for _, host := range hosts {
			nagiosxi.AddHost(nagiosxiConfig, host, force)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	addCmd.AddCommand(addHostsCmd)
}
