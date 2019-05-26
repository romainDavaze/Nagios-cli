package nagiosxi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Command represents the NagiosXI command object
type Command struct {
	Name        string `schema:"command_name" yaml:"name"`
	CommandLine string `schema:"command_line" yaml:"commandLine"`
}

// AddCommand adds a command to NagiosXI
func AddCommand(config Config, command Command) {
	values := make(map[string][]string)

	encoder := InitEncoder()
	err := encoder.Encode(command, values)
	if err != nil {
		log.Fatalf("Error while encoding command %q: %s", command.Name, err)
	}

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+"/"+config.BasePath+"/config/command?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		log.Fatalf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding command %q:\n%s", command.Name, string(body))
}

// DeleteCommand deletes a commands from NagiosXI
func DeleteCommand(config Config, command Command) {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + "/" + config.BasePath + "/config/command?apikey=" +
		config.APIKey + "&pretty=1&command_name=" + command.Name)
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Deleting command %q:\n%s", command.Name, string(body))
}

// ParseCommands parses NagiosXI commands from a given yaml file
func ParseCommands(file string) []Command {
	var objects map[string][]map[string]interface{}

	content, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(content, &objects)

	obj := objects["commands"]
	if len(obj) == 0 {
		log.Fatal("There is no commands object in the given file")
	}

	var commands []Command
	mapstructure.Decode(obj, &commands)

	return commands
}
