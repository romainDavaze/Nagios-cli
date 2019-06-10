package cmd

import (
	"log"

	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
)

var deleteContactsCmd = &cobra.Command{
	Use:   "contacts <file>",
	Short: "Delete NagiosXI contacts",
	Long:  "Delete NagiosXI contacts",
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		contacts, err := nagiosxi.ParseContacts(args[0])
		if err != nil {
			log.Fatal(err)
		}

		for _, contact := range contacts {
			err := nagiosxi.DeleteContact(nagiosxiConfig, contact)
			if err != nil {
				log.Fatal(err)
			}
		}

		if applyConfig {
			err := nagiosxi.ApplyConfig(nagiosxiConfig)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteContactsCmd)
}
