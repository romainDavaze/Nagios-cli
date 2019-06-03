package nagiosxi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Service represents the NagiosXI service object
type Service struct {
	CheckCommand         string   `schema:"check_command,omitempty" yaml:"checkCommand"`
	CheckCommandArgs     []string `schema:"-" yaml:"checkCommandArgs"`
	CheckInterval        string   `schema:"check_interval,omitempty" yaml:"checkInterval"`
	CheckPeriod          string   `schema:"check_period,omitempty" yaml:"checkPeriod"`
	ConfigName           string   `schema:"config_name,omitempty" yaml:"configName"`
	Contacts             []string `schema:"contacts,omitempty" yaml:"contacts"`
	ContactGroups        []string `schema:"contact_groups,omitempty" yaml:"contactGroups"`
	DisplayName          string   `schema:"display_name,omitempty" yaml:"displayName"`
	Hosts                []string `schema:"host_name,omitempty" yaml:"hosts"`
	HostGroups           []string `schema:"hostgroup_name,omitempty" yaml:"hostGroups"`
	MaxCheckAttempts     string   `schema:"max_check_attempts,omitempty" yaml:"maxCheckAttempts"`
	NotificationInterval string   `schema:"notification_interval,omitempty" yaml:"notificationInterval"`
	NotificationPeriod   string   `schema:"notification_period,omitempty" yaml:"notificationPeriod"`
	RetryInterval        string   `schema:"retry_interval,omitempty" yaml:"retryInterval"`
	ServiceDescription   string   `schema:"service_description,omitempty" yaml:"serviceDescription"`
	ServiceGroups        []string `schema:"servicegroups,omitempty" yaml:"serviceGroups"`
	Templates            []string `schema:"use,omitempty" yaml:"templates"`
}

// Encode encodes service into a map[string][]string
func (service *Service) Encode(force bool) (map[string][]string, error) {
	var argsString string
	values := make(map[string][]string)

	encoder := InitEncoder()
	err := encoder.Encode(service, values)
	if err != nil {
		log.Fatalf("Error while encoding service %q for hosts [%q]: %s", service.ServiceDescription, strings.Join(service.Hosts, ","), err)
	}

	for _, arg := range service.CheckCommandArgs {
		argsString += "\\!" + arg
	}
	values["check_command"] = []string{service.CheckCommand + argsString}

	if len(values["config_name"]) == 1 {
		values["config_name"] = []string{strings.Join(service.Hosts, " ")}
	}

	values["force"] = []string{BoolToStr(force)}

	return values, err
}

// AddService adds a service to NagiosXI
func AddService(config Config, service Service, force bool) {
	values, _ := service.Encode(force)

	fmt.Println(values)

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
