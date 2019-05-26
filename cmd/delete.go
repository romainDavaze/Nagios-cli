package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete NagiosXI objects",
	Long:  "Delete NagiosXI objects",
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.PersistentFlags().BoolVar(&applyConfig, "applyconfig", false, "indicate whether changes should be applied or not (false by default)")
	deleteCmd.PersistentFlags().StringVarP(&objectsFile, "file", "f", "", "file containing services to add")
	cobra.MarkFlagRequired(deleteCmd.PersistentFlags(), "file")
}
