package cmd

import (
	"log"

	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var addHostsCmd = &cobra.Command{
	Use:   "hosts <file>",
	Short: "Add NagiosXI hosts",
	Long:  "Add NagiosXI hosts",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		hosts, err := nagiosxi.ParseHosts(args[0])
		if err != nil {
			log.Fatal(err)
		}

		for _, host := range hosts {
			err := nagiosxi.AddHost(nagiosxiConfig, host, force)
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
	addCmd.AddCommand(addHostsCmd)
}
