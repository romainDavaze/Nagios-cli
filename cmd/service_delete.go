package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a single or multiples Nagios service(s)",
	Long:  "Delete a single or multiples Nagios service(s)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	serviceCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&nagiosFile, "file", "f", "", "file containing services to delete")
	cobra.MarkFlagRequired(deleteCmd.Flags(), "file")
}
