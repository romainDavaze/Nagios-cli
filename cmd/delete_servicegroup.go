package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteServicegroupsCmd = &cobra.Command{
	Use:   "servicegroups",
	Short: "Delete NagiosXI servicegroups",
	Long:  "Delete NagiosXI servicegroups",
	Run: func(cmd *cobra.Command, args []string) {
		servicegroups := nagiosxi.ParseServicegroups(objectsFile)

		for _, servicegroup := range servicegroups {
			nagiosxi.DeleteServicegroup(nagiosxiConfig, servicegroup)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteServicegroupsCmd)
}
