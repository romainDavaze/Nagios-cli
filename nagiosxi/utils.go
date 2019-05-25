package nagiosxi

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Config holding NagiosXI configuration
type Config struct {
	APIKey   string `yaml:"apiKey"`
	BasePath string `yaml:"basePath"`
	Host     string `yaml:"nagiosxiHost"`
	Protocol string `yaml:"protocol"`
}

// AddApplyConfigToJSON adds the applyconfig attribute to a json string
func AddApplyConfigToJSON(data []byte) ([]byte, error) {
	var d map[string]interface{}
	json.Unmarshal(data, &d)
	d["applyconfig"] = 0
	return json.Marshal(d)
}

// ApplyConfig applies changes made previously. It also asks for user confirmation to make sure the user wants to do it.
func ApplyConfig(config Config) {
	var choice string
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\n\nAre you sure you want to apply current NagiosXI configuration [y/N] ? ")
	choice, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	choice = strings.ToLower(strings.TrimSpace(choice))

	if choice == "y" || choice == "yes" {
		resp, err := http.Get(config.Protocol + "://" + config.Host + "/" + config.BasePath + "/system/applyconfig?apikey=" + config.APIKey)
		if err != nil {
			log.Fatalf("Error while making POST request to NagiosXI API: %s", err)
		}

		defer resp.Body.Close()

		fmt.Println("Configuration applied !")
	} else {
		fmt.Println("Not applying configuration.")
	}
}
