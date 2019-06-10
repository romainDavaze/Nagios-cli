package nagiosxi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Host represents the NagiosXI host object
type Host struct {
	Address              string   `json:"address" schema:"address,omitempty" yaml:"address"`
	Alias                string   `json:"alias" schema:"alias,omitempty" yaml:"alias"`
	CheckCommand         string   `json:"check_command" schema:"check_command,omitempty" yaml:"checkCommand"`
	CheckCommandArgs     []string `json:"-" schema:"-" yaml:"checkCommandArgs"`
	CheckInterval        string   `json:"check_interval" schema:"check_interval,omitempty" yaml:"checkInterval"`
	CheckPeriod          string   `json:"check_period" schema:"check_period,omitempty" yaml:"checkPeriod"`
	Contacts             []string `json:"contacts" schema:"contacts,omitempty" yaml:"contacts"`
	ContactGroups        []string `json:"contact_groups" schema:"contact_groups,omitempty" yaml:"contactGroups"`
	DisplayName          string   `json:"display_name" schema:"display_name,omitempty" yaml:"displayName"`
	HostGroups           []string `json:"hostgroups" schema:"hostgroups,omitempty" yaml:"hostGroups"`
	MaxCheckAttempts     string   `json:"max_check_attempts" schema:"max_check_attempts,omitempty" yaml:"maxCheckAttempts"`
	Name                 string   `json:"host_name" schema:"host_name,omitempty" yaml:"name"`
	NotificationInterval string   `json:"notification_interval" schema:"notification_interval,omitempty" yaml:"notificationInterval"`
	NotificationPeriod   string   `json:"notification_period" schema:"notification_period,omitempty" yaml:"notificationPeriod"`
	Parents              []string `json:"parents" schema:"parents,omitempty" yaml:"parents"`
	RetryInterval        string   `json:"retry_interval" schema:"retry_interval,omitempty" yaml:"retryInterval"`
	Templates            []string `json:"use" schema:"use,omitempty" yaml:"templates"`
}

// Encode encodes service into a map[string][]string
func (host *Host) Encode() (map[string][]string, error) {
	var argsString string
	values := make(map[string][]string)

	encoder := InitEncoder()
	err := encoder.Encode(host, values)
	if err != nil {
		log.Fatalf("Error while encoding host %q: %s", host.Name, err)
	}

	for _, arg := range host.CheckCommandArgs {
		argsString += "\\!" + arg
	}
	values["check_command"] = []string{host.CheckCommand + argsString}

	return values, err
}

// AddHost adds a host to NagiosXI
func AddHost(config Config, host Host, force bool) error {
	values, _ := host.Encode()

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+":"+strconv.Itoa(int(config.Port))+"/"+config.BasePath+"/config/host?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		return fmt.Errorf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding host %q:\n%s", host.Name, string(body))

	return nil
}

// DeleteHost deletes a hosts from NagiosXI
func DeleteHost(config Config, host Host) error {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/host?apikey=" +
		config.APIKey + "&pretty=1&host_name=" + host.Name)
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Deleting host %q:\n%s", host.Name, string(body))

	return nil
}

// GetHost retrives a host from NagiosXI
func GetHost(config Config, hostName string) (Host, error) {
	hosts := []Host{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/host?apikey=" +
		config.APIKey + "&pretty=1&host_name=" + hostName)
	resp, err := http.Get(fullURL)
	if err != nil {
		return Host{}, fmt.Errorf("Error while retrieving %s host from NagiosXI: %s", hostName, err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &hosts)
	if err != nil {
		return Host{}, fmt.Errorf("Error while unmarshalling %s host from NagiosXI: %s", hostName, err)
	}

	if len(hosts) == 0 {
		return Host{}, fmt.Errorf("Could not retrieve host %s from NagiosXI", hostName)
	}

	return hosts[0], nil
}

// ParseHosts parses NagiosXI hosts from a given yaml file
func ParseHosts(file string) ([]Host, error) {
	hosts := []Host{}
	var objects map[string][]map[string]interface{}

	content, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(content, &objects)

	obj := objects["hosts"]
	if len(obj) == 0 {
		return hosts, fmt.Errorf("There is no hosts object in the given file")
	}

	mapstructure.Decode(obj, &hosts)

	return hosts, nil
}
