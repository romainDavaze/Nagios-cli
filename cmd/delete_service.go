package cmd

import (
	"log"

	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteServicesCmd = &cobra.Command{
	Use:   "services <file>",
	Short: "Delete NagiosXI services",
	Long:  "Delete NagiosXI services",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		services, err := nagiosxi.ParseServices(args[0])
		if err != nil {
			log.Fatal(err)
		}

		for _, service := range services {
			err := nagiosxi.DeleteService(nagiosxiConfig, service)
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
	deleteCmd.AddCommand(deleteServicesCmd)
}
