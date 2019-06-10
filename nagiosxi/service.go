package nagiosxi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Service represents the NagiosXI service object
type Service struct {
	CheckCommand         string   `json:"check_command" schema:"check_command,omitempty" yaml:"checkCommand"`
	CheckCommandArgs     []string `json:"-" schema:"-" yaml:"checkCommandArgs"`
	CheckInterval        string   `json:"check_interval" schema:"check_interval,omitempty" yaml:"checkInterval"`
	CheckPeriod          string   `json:"check_period" schema:"check_period,omitempty" yaml:"checkPeriod"`
	ConfigName           string   `json:"config_name" schema:"config_name,omitempty" yaml:"configName"`
	Contacts             []string `json:"Services" schema:"Services,omitempty" yaml:"Services"`
	ContactGroups        []string `json:"Service_groups" schema:"Service_groups,omitempty" yaml:"ServiceGroups"`
	DisplayName          string   `json:"display_name" schema:"display_name,omitempty" yaml:"displayName"`
	Hosts                []string `json:"host_name" schema:"host_name,omitempty" yaml:"hosts"`
	HostGroups           []string `json:"hostgroup_name" schema:"hostgroup_name,omitempty" yaml:"hostGroups"`
	MaxCheckAttempts     string   `json:"max_check_attempts" schema:"max_check_attempts,omitempty" yaml:"maxCheckAttempts"`
	NotificationInterval string   `json:"notification_interval" schema:"notification_interval,omitempty" yaml:"notificationInterval"`
	NotificationPeriod   string   `json:"notification_period" schema:"notification_period,omitempty" yaml:"notificationPeriod"`
	RetryInterval        string   `json:"retry_interval" schema:"retry_interval,omitempty" yaml:"retryInterval"`
	ServiceDescription   string   `json:"service_description" schema:"service_description,omitempty" yaml:"serviceDescription"`
	ServiceGroups        []string `json:"servicegroups" schema:"servicegroups,omitempty" yaml:"serviceGroups"`
	Templates            []string `json:"use" schema:"use,omitempty" yaml:"templates"`
}

// Encode encodes a service into a map[string][]string
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
func AddService(config Config, service Service, force bool) error {
	values, err := service.Encode(force)
	if err != nil {
		return fmt.Errorf("Error while encoding service %q for hosts [%s]: %s", service.ServiceDescription, strings.Join(service.Hosts, ","), err)
	}

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+":"+strconv.Itoa(int(config.Port))+"/"+config.BasePath+"/config/service?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		return fmt.Errorf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding service %q for host(s) [%s]:\n%s", service.ServiceDescription, strings.Join(service.Hosts, ","), string(body))

	return nil
}

// DeleteService deletes a service from NagiosXI
func DeleteService(config Config, service Service) error {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/service?apikey=" +
		config.APIKey + "&pretty=1&" + EncodeStringArrayForDeletion(service.Hosts, "host_name") + "&service_description=" +
		url.QueryEscape(service.ServiceDescription))
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Deleting service %q for host(s) [%s]:\n%s", service.ServiceDescription, strings.Join(service.Hosts, ","), string(body))

	return nil
}

// GetService retrives a service from NagiosXI
func GetService(config Config, configName, serviceDescription string) (Service, error) {
	services := []Service{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/service?apikey=" +
		config.APIKey + "&pretty=1&config_name=" + url.QueryEscape(configName) + "&service_description=" + url.QueryEscape(serviceDescription))
	resp, err := http.Get(fullURL)
	if err != nil {
		return Service{}, fmt.Errorf("Error while retrieving %q service for host(s) [%s] from NagiosXI: %s", serviceDescription, configName, err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &services)
	if err != nil {
		return Service{}, fmt.Errorf("Error while unmarshalling %q service for host(s) [%s] from NagiosXI: %s", serviceDescription, configName, err)
	}

	if len(services) == 0 {
		return Service{}, fmt.Errorf("Could not retrieve service %q for host(s) [%s] from NagiosXI", serviceDescription, configName)
	}

	return services[0], nil
}

// ParseServices parses NagiosXI services from a given yaml file
func ParseServices(file string) ([]Service, error) {
	var objects map[string][]map[string]interface{}
	services := []Service{}

	content, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(content, &objects)

	obj := objects["services"]
	if len(obj) == 0 {
		return services, fmt.Errorf("There is no services object in the given file")
	}

	mapstructure.Decode(obj, &services)

	return services, nil
}
