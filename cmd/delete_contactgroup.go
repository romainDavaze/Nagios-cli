package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteContactgroupsCmd = &cobra.Command{
	Use:   "contactgroups",
	Short: "Delete NagiosXI contactgroups",
	Long:  "Delete NagiosXI contactgroups",
	Run: func(cmd *cobra.Command, args []string) {
		contactgroups := nagiosxi.ParseContactgroups(objectsFile)

		for _, contactgroup := range contactgroups {
			nagiosxi.DeleteContactgroup(nagiosxiConfig, contactgroup)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteContactgroupsCmd)
}
