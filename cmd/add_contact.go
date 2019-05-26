package cmd

import (
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var addContactsCmd = &cobra.Command{
	Use:   "contacts",
	Short: "Add NagiosXI contacts",
	Long:  "Add NagiosXI contacts",
	Run: func(cmd *cobra.Command, args []string) {
		contacts := nagiosxi.ParseContacts(objectsFile)

		for _, contact := range contacts {
			nagiosxi.AddContact(nagiosxiConfig, contact)
		}

		if applyConfig {
			nagiosxi.ApplyConfig(nagiosxiConfig)
		}
	},
}

func init() {
	addCmd.AddCommand(addContactsCmd)
}
