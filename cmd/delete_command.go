package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteCommandCmd = &cobra.Command{
	Use:   "commands <file>",
	Short: "Delete NagiosXI commands",
	Long:  "Delete NagiosXI commands",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		commands := nagiosxi.ParseCommands(objectsFile)

		for _, command := range commands {
			nagiosxi.DeleteCommand(nagiosxiConfig, command)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteCommandCmd)
}
