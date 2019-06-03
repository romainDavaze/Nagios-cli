package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add NagiosXI objects",
	Long:  "Add NagiosXI objects",
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().BoolVar(&applyConfig, "applyconfig", false, "indicate whether changes should be applied or not (false by default)")
	addCmd.PersistentFlags().BoolVar(&force, "force", false, "force API call")
	addCmd.PersistentFlags().StringVarP(&objectsFile, "file", "f", "", "file containing NagiosXI hosts")
	cobra.MarkFlagRequired(addCmd.PersistentFlags(), "file")
}
