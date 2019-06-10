package cmd

import (
	"log"

	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteHostgroupsCmd = &cobra.Command{
	Use:   "hostgroups <file>",
	Short: "Delete NagiosXI hostgroups",
	Long:  "Delete NagiosXI hostgroups",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		hostgroups, err := nagiosxi.ParseHostgroups(args[0])
		if err != nil {
			log.Fatal(err)
		}

		for _, hostgroup := range hostgroups {
			err := nagiosxi.DeleteHostgroup(nagiosxiConfig, hostgroup)
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
	deleteCmd.AddCommand(deleteHostgroupsCmd)
}
