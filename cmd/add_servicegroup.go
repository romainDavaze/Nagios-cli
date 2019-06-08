package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var addServicegroupsCmd = &cobra.Command{
	Use:   "servicegroups <file>",
	Short: "Add NagiosXI servicegroups",
	Long:  "Add NagiosXI servicegroups",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		servicegroups := nagiosxi.ParseServicegroups(objectsFile)

		for _, servicegroup := range servicegroups {
			nagiosxi.AddServicegroup(nagiosxiConfig, servicegroup, force)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	addCmd.AddCommand(addServicegroupsCmd)
}
