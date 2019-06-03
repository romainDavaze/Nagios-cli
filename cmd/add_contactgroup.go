package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var addContactgroupsCmd = &cobra.Command{
	Use:   "contactgroups",
	Short: "Add NagiosXI contactgroups",
	Long:  "Add NagiosXI contactgroups",
	Run: func(cmd *cobra.Command, args []string) {
		contactgroups := nagiosxi.ParseContactgroups(objectsFile)

		for _, contactgroup := range contactgroups {
			nagiosxi.AddContactgroup(nagiosxiConfig, contactgroup, force)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	addCmd.AddCommand(addContactgroupsCmd)
}
