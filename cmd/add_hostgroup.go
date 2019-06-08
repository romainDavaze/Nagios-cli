package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var addHostgroupsCmd = &cobra.Command{
	Use:   "hostgroups <file>",
	Short: "Add NagiosXI hostgroups",
	Long:  "Add NagiosXI hostgroups",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		hostgroups := nagiosxi.ParseHostgroups(objectsFile)

		for _, hostgroup := range hostgroups {
			nagiosxi.AddHostgroup(nagiosxiConfig, hostgroup, force)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	addCmd.AddCommand(addHostgroupsCmd)
}
