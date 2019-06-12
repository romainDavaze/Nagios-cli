package cmd

import (
	"log"

	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var applyConfigCmd = &cobra.Command{
	Use:   "applyconfig",
	Short: "Apply current NagiosXI configuration",
	Long:  "Apply current NagiosXI configuration",
	Run: func(cmd *cobra.Command, args []string) {
		err := nagiosxi.ApplyConfig(nagiosxiConfig)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(applyConfigCmd)
}
