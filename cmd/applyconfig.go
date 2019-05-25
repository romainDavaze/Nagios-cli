package cmd

import (
	"github.com/romainDavaze/nagios-cli/nagios"
	"github.com/spf13/cobra"
)

var applyConfigCmd = &cobra.Command{
	Use:   "applyConfig",
	Short: "Applies current Nagios configuration",
	Long:  "Applies current Nagios configuration",
	Run: func(cmd *cobra.Command, args []string) {
		nagios.ApplyConfig(nagiosConfig)
	},
}

func init() {
	rootCmd.AddCommand(applyConfigCmd)
}
