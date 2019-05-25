package cmd

import (
	"io/ioutil"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Interacts with NagiosXI host object",
	Long:  "Interacts with NagiosXI host object",
}

func init() {
	rootCmd.AddCommand(hostCmd)

	hostCmd.PersistentFlags().BoolVar(&applyConfig, "applyconfig", false, "indicate whether changes should be applied or not (false by default)")
	hostCmd.PersistentFlags().StringVarP(&nagiosxiFile, "file", "f", "", "file containing NagiosXI hosts")
	cobra.MarkFlagRequired(hostCmd.PersistentFlags(), "file")
}

func parseHosts() []nagiosxi.Host {
	var objects map[string][]map[string]interface{}

	content, _ := ioutil.ReadFile(nagiosxiFile)
	yaml.Unmarshal(content, &objects)

	obj := objects["hosts"]
	if len(obj) == 0 {
		log.Fatal("There is no hosts object in the given file")
	}

	var hosts []nagiosxi.Host
	mapstructure.Decode(obj, &hosts)

	return hosts
}
