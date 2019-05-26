package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var addCommandCmd = &cobra.Command{
	Use:   "commands",
	Short: "Add NagiosXI commands",
	Long:  "Add NagiosXI commands",
	Run: func(cmd *cobra.Command, args []string) {
		commands := nagiosxi.ParseCommands(objectsFile)

		for _, command := range commands {
			nagiosxi.AddCommand(nagiosxiConfig, command)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	addCmd.AddCommand(addCommandCmd)
}
