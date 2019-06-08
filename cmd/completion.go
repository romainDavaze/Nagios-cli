package cmd

import (
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates completion scripts for shells",
	Long:  `Generates completion scripts for shells`,
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
