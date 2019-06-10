package cmd

import (
	"log"

	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteContactgroupsCmd = &cobra.Command{
	Use:   "contactgroups <file>",
	Short: "Delete NagiosXI contactgroups",
	Long:  "Delete NagiosXI contactgroups",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		contactgroups, err := nagiosxi.ParseContactgroups(args[0])
		if err != nil {
			log.Fatal(err)
		}

		for _, contactgroup := range contactgroups {
			err := nagiosxi.DeleteContactgroup(nagiosxiConfig, contactgroup)
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
	deleteCmd.AddCommand(deleteContactgroupsCmd)
}
