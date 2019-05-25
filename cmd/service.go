package cmd

import (
	"io/ioutil"
	"log"
	"nagios-cli/nagios"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Interacts with Nagios service object",
	Long:  "Interacts with Nagios service object",
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}

func parseServices() []nagios.Service {
	var objects map[string][]map[string]interface{}

	content, _ := ioutil.ReadFile(nagiosFile)
	yaml.Unmarshal(content, &objects)

	obj := objects["services"]
	if len(obj) == 0 {
		log.Fatal("There is no services object in the given file")
	}

	var services []nagios.Service
	mapstructure.Decode(obj, &services)

	return services
}
