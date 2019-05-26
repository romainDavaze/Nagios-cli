package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var applyConfigCmd = &cobra.Command{
	Use:   "applyConfig",
	Short: "Apply current NagiosXI configuration",
	Long:  "Apply current NagiosXI configuration",
	Run: func(cmd *cobra.Command, args []string) {
		nagiosxi.ApplyConfig(nagiosxiConfig)
	},
}

func init() {
	rootCmd.AddCommand(applyConfigCmd)
}
