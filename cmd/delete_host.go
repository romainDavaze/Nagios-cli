package cmd

import (
	"log"

	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteHostsCmd = &cobra.Command{
	Use:   "hosts <file>",
	Short: "Delete NagiosXI hosts",
	Long:  "Delete NagiosXI hosts",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		hosts, err := nagiosxi.ParseHosts(args[0])
		if err != nil {
			log.Fatal(err)
		}

		for _, host := range hosts {
			err := nagiosxi.DeleteHost(nagiosxiConfig, host)
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
	deleteCmd.AddCommand(deleteHostsCmd)
}
