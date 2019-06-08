package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var addCommandCmd = &cobra.Command{
	Use:   "commands <file>",
	Short: "Add NagiosXI commands",
	Long:  "Add NagiosXI commands",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		commands := nagiosxi.ParseCommands(objectsFile)

		for _, command := range commands {
			nagiosxi.AddCommand(nagiosxiConfig, command, force)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	addCmd.AddCommand(addCommandCmd)
}
