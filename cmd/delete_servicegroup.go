package cmd

import (
	"log"

	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteServicegroupsCmd = &cobra.Command{
	Use:   "servicegroups <file>",
	Short: "Delete NagiosXI servicegroups",
	Long:  "Delete NagiosXI servicegroups",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		servicegroups, err := nagiosxi.ParseServicegroups(args[0])
		if err != nil {
			log.Fatal(err)
		}

		for _, servicegroup := range servicegroups {
			err := nagiosxi.DeleteServicegroup(nagiosxiConfig, servicegroup)
			if err != nil {
				log.Fatal(err)
			}
		}

		if applyConfig {
			err := nagiosxi.ApplyConfig(nagiosxiConfig)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteServicegroupsCmd)
}
