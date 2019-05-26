package nagiosxi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/schema"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Service represents the NagiosXI service object
type Service struct {
	CheckCommand         string   `schema:"check_command,omitempty" yaml:"checkCommand"`
	CheckCommandArgs     []string `schema:"-" yaml:"checkCommandArgs"`
	CheckInterval        string   `schema:"check_interval" yaml:"checkInterval"`
	CheckPeriod          string   `schema:"check_period" yaml:"checkPeriod"`
	ConfigName           string   `schema:"config_name" yaml:"configName"`
	Contacts             []string `schema:"contacts,omitempty" yaml:"contacts"`
	ContactGroups        []string `schema:"contact_groups,omitempty" yaml:"contactGroups"`
	DisplayName          string   `schema:"display_name" yaml:"displayName"`
	Hosts                []string `schema:"host_name" yaml:"hosts"`
	HostGroups           []string `schema:"hostgroup_name" yaml:"hostGroups"`
	MaxCheckAttempts     string   `schema:"max_check_attempts" yaml:"maxCheckAttempts"`
	NotificationInterval string   `schema:"notification_interval" yaml:"notificationInterval"`
	NotificationPeriod   string   `schema:"notification_period" yaml:"notificationPeriod"`
	RetryInterval        string   `schema:"retry_interval" yaml:"retryInterval"`
	ServiceDescription   string   `schema:"service_description" yaml:"serviceDescription"`
	ServiceGroups        []string `schema:"servicegroups" yaml:"serviceGroups"`
	Templates            []string `schema:"use" yaml:"templates"`
}

// Encode encodes service into a map[string][]string
func (s *Service) Encode() (map[string][]string, error) {
	var argsString string
	values := make(map[string][]string)

	encoder := schema.NewEncoder()
	encoder.RegisterEncoder([]string{}, EncodeStringArray)

	err := encoder.Encode(s, values)

	for _, arg := range s.CheckCommandArgs {
		argsString += "\\!" + arg
	}
	values["check_command"] = []string{s.CheckCommand + argsString}

	if len(values["config_name"]) == 1 {
		values["config_name"] = []string{strings.Join(s.Hosts, " ")}
	}

	return values, err
}

// AddService adds a service to NagiosXI
func AddService(config Config, service Service) {
	values, _ := service.Encode()

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+"/"+config.BasePath+"/config/service?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		log.Fatalf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding service %q for host(s) [%s]:\n%s", service.ServiceDescription, strings.Join(service.Hosts, ","), string(body))
}

// DeleteService deletes a service from NagiosXI
func DeleteService(config Config, service Service) {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + "/" + config.BasePath + "/config/service?apikey=" +
		config.APIKey + "&pretty=1&" + EncodeStringArrayForDeletion(service.Hosts, "host_name") + "&service_description=" +
		url.QueryEscape(service.ServiceDescription))
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Deleting service %q for host(s) [%s]:\n%s", service.ServiceDescription, strings.Join(service.Hosts, ","), string(body))
}

// ParseServices parses NagiosXI services from a given yaml file
func ParseServices(file string) []Service {
	var objects map[string][]map[string]interface{}

	content, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(content, &objects)

	obj := objects["services"]
	if len(obj) == 0 {
		log.Fatal("There is no services object in the given file")
	}

	var services []Service
	mapstructure.Decode(obj, &services)

	return services
}
