package nagiosxi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Host represents the NagiosXI host object
type Host struct {
	Address              string   `schema:"address" yaml:"address"`
	Alias                string   `schema:"alias" yaml:"alias"`
	CheckCommand         string   `schema:"check_command,omitempty" yaml:"checkCommand"`
	CheckCommandArgs     []string `schema:"-" yaml:"checkCommandArgs"`
	CheckInterval        string   `schema:"check_interval" yaml:"checkInterval"`
	CheckPeriod          string   `schema:"check_period" yaml:"checkPeriod"`
	Contacts             []string `schema:"contacts,omitempty" yaml:"contacts"`
	ContactGroups        []string `schema:"contact_groups,omitempty" yaml:"contactGroups"`
	DisplayName          string   `schema:"display_name" yaml:"displayName"`
	HostGroups           []string `schema:"hostgroups" yaml:"hostGroups"`
	MaxCheckAttempts     string   `schema:"max_check_attempts" yaml:"maxCheckAttempts"`
	Name                 string   `schema:"host_name" yaml:"name"`
	NotificationInterval string   `schema:"notification_interval" yaml:"notificationInterval"`
	NotificationPeriod   string   `schema:"notification_period" yaml:"notificationPeriod"`
	Parents              []string `schema:"parents" yaml:"parents"`
	RetryInterval        string   `schema:"retry_interval" yaml:"retryInterval"`
	Templates            []string `schema:"use" yaml:"templates"`
}

// Encode encodes service into a map[string][]string
func (h *Host) Encode() (map[string][]string, error) {
	var argsString string
	values := make(map[string][]string)

	encoder := schema.NewEncoder()
	encoder.RegisterEncoder([]string{}, EncodeStringArray)

	err := encoder.Encode(h, values)

	for _, arg := range h.CheckCommandArgs {
		argsString += "\\!" + arg
	}
	values["check_command"] = []string{h.CheckCommand + argsString}

	return values, err
}

// AddHost adds a host to NagiosXI
func AddHost(config Config, host Host) {
	values, _ := host.Encode()

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+"/"+config.BasePath+"/config/host?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		log.Fatalf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding host %q:\n%s", host.Name, string(body))
}

// DeleteHost deletes a hosts from NagiosXI
func DeleteHost(config Config, host Host) {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + "/" + config.BasePath + "/config/host?apikey=" +
		config.APIKey + "&pretty=1&host_name=" + host.Name)
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Deleting host %q:\n%s", host.Name, string(body))
}

// ParseHosts parses NagiosXI hosts from a given yaml file
func ParseHosts(file string) []Host {
	var objects map[string][]map[string]interface{}

	content, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(content, &objects)

	obj := objects["hosts"]
	if len(obj) == 0 {
		log.Fatal("There is no hosts object in the given file")
	}

	var hosts []Host
	mapstructure.Decode(obj, &hosts)

	return hosts
}
