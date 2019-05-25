package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var applyConfigCmd = &cobra.Command{
	Use:   "applyConfig",
	Short: "Applies current NagiosXI configuration",
	Long:  "Applies current NagiosXI configuration",
	Run: func(cmd *cobra.Command, args []string) {
		nagiosxi.ApplyConfig(nagiosxiConfig)
	},
}

func init() {
	rootCmd.AddCommand(applyConfigCmd)
}
