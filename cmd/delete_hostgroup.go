package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteHostgroupsCmd = &cobra.Command{
	Use:   "hostgroups",
	Short: "Delete NagiosXI hostgroups",
	Long:  "Delete NagiosXI hostgroups",
	Run: func(cmd *cobra.Command, args []string) {
		hostgroups := nagiosxi.ParseHostgroups(objectsFile)

		for _, hostgroup := range hostgroups {
			nagiosxi.DeleteHostgroup(nagiosxiConfig, hostgroup)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteHostgroupsCmd)
}
