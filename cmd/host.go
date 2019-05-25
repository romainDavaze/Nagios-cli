package cmd

import (
	"io/ioutil"
	"log"
	"nagios-cli/nagios"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Interacts with Nagios host object",
	Long:  "Interacts with Nagios host object",
}

func init() {
	rootCmd.AddCommand(hostCmd)

	hostCmd.PersistentFlags().BoolVar(&applyConfig, "applyconfig", false, "indicate whether changes should be applied or not (false by default)")
	hostCmd.PersistentFlags().StringVarP(&nagiosFile, "file", "f", "", "file containing Nagios hosts")
	cobra.MarkFlagRequired(hostCmd.PersistentFlags(), "file")
}

func parseHosts() []nagios.Host {
	var objects map[string][]map[string]interface{}

	content, _ := ioutil.ReadFile(nagiosFile)
	yaml.Unmarshal(content, &objects)

	obj := objects["hosts"]
	if len(obj) == 0 {
		log.Fatal("There is no hosts object in the given file")
	}

	var hosts []nagios.Host
	mapstructure.Decode(obj, &hosts)

	return hosts
}
