package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteHostsCmd = &cobra.Command{
	Use:   "hosts <file>",
	Short: "Delete NagiosXI hosts",
	Long:  "Delete NagiosXI hosts",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		hosts := nagiosxi.ParseHosts(objectsFile)

		for _, host := range hosts {
			nagiosxi.DeleteHost(nagiosxiConfig, host)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteHostsCmd)
}
