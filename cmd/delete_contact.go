package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteContactsCmd = &cobra.Command{
	Use:   "contacts",
	Short: "Delete NagiosXI contacts",
	Long:  "Delete NagiosXI contacts",
	Run: func(cmd *cobra.Command, args []string) {
		contacts := nagiosxi.ParseContacts(objectsFile)

		for _, contact := range contacts {
			nagiosxi.DeleteContact(nagiosxiConfig, contact)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteContactsCmd)
}
