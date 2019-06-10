package cmd

import (
	"log"

	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var addContactgroupsCmd = &cobra.Command{
	Use:   "contactgroups <file>",
	Short: "Add NagiosXI contactgroups",
	Long:  "Add NagiosXI contactgroups",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		contactgroups, err := nagiosxi.ParseContactgroups(args[0])
		if err != nil {
			log.Fatal(err)
		}

		for _, contactgroup := range contactgroups {
			err := nagiosxi.AddContactgroup(nagiosxiConfig, contactgroup, force)
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
	addCmd.AddCommand(addContactgroupsCmd)
}
