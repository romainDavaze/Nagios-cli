package nagiosxi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Command represents the NagiosXI command object
type Command struct {
	Name        string `json:"command_name" schema:"command_name,omitempty" yaml:"name"`
	CommandLine string `json:"command_line" schema:"command_line,omitempty" yaml:"commandLine"`
}

// AddCommand adds a command to NagiosXI
func AddCommand(config Config, command Command, force bool) error {
	values := make(map[string][]string)

	encoder := InitEncoder()
	err := encoder.Encode(command, values)
	if err != nil {
		return fmt.Errorf("Error while encoding command %q: %s", command.Name, err)
	}

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+":"+strconv.Itoa(int(config.Port))+"/"+config.BasePath+"/config/command?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		return fmt.Errorf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding command %q:\n%s", command.Name, string(body))

	return nil
}

// DeleteCommand deletes a commands from NagiosXI
func DeleteCommand(config Config, command Command) error {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/command?apikey=" +
		config.APIKey + "&pretty=1&command_name=" + command.Name)
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Deleting command %q:\n%s", command.Name, string(body))

	return nil
}

// GetCommand retrives a command from NagiosXI
func GetCommand(config Config, commandName string) (Command, error) {
	commands := []Command{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/command?apikey=" +
		config.APIKey + "&pretty=1&command_name=" + commandName)
	resp, err := http.Get(fullURL)
	if err != nil {
		return Command{}, fmt.Errorf("Error while retrieving %s command from NagiosXI: %s", commandName, err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &commands)
	if err != nil {
		return Command{}, fmt.Errorf("Error while unmarshalling %s command from NagiosXI: %s", commandName, err)
	}

	if len(commands) == 0 {
		return Command{}, fmt.Errorf("Could not retrieve command %s from NagiosXI", commandName)
	}

	return commands[0], nil
}

// ParseCommands parses NagiosXI commands from a given yaml file
func ParseCommands(file string) ([]Command, error) {
	commands := []Command{}
	var objects map[string][]map[string]interface{}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return commands, fmt.Errorf("Error while reading objects file: %s", err)
	}

	yaml.Unmarshal(content, &objects)

	obj := objects["commands"]
	if len(obj) == 0 {
		return commands, fmt.Errorf("There is no commands object in the given file")
	}

	mapstructure.Decode(obj, &commands)

	return commands, nil
}
