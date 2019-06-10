package cmd

import (
	"log"

	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteCommandCmd = &cobra.Command{
	Use:   "commands <file>",
	Short: "Delete NagiosXI commands",
	Long:  "Delete NagiosXI commands",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		commands, err := nagiosxi.ParseCommands(args[0])
		if err != nil {
			log.Fatal(err)
		}

		for _, command := range commands {
			err := nagiosxi.DeleteCommand(nagiosxiConfig, command)
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
	deleteCmd.AddCommand(deleteCommandCmd)
}
