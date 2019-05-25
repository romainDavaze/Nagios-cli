package nagiosxi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
)

// Service represents the NagiosXI service object
type Service struct {
	CheckCommand         string   `schema:"check_command,omitempty"`
	CheckCommandArgs     []string `schema:"-"`
	CheckInterval        string   `schema:"check_interval"`
	CheckPeriod          string   `schema:"check_period"`
	Contacts             []string `schema:"contacts,omitempty"`
	ContactGroups        []string `schema:"contact_groups,omitempty"`
	HostName             string   `schema:"host_name"`
	MaxCheckAttempts     string   `schema:"max_check_attempts"`
	NotificationInterval string   `schema:"notification_interval"`
	NotificationPeriod   string   `schema:"notification_period"`
	RetryInterval        string   `schema:"retry_interval"`
	ServiceDescription   string   `schema:"service_description"`
}

// Encode encodes service into a map[string][]string
func (s *Service) Encode() (map[string][]string, error) {
	var argsString string
	values := make(map[string][]string)
	encoder := schema.NewEncoder()

	err := encoder.Encode(s, values)

	for _, arg := range s.CheckCommandArgs {
		argsString += "\\!" + arg
	}
	values["check_command"] = []string{s.CheckCommand + argsString}

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
	fmt.Printf("Adding service %q for host %q:\n%s", service.ServiceDescription, service.HostName, string(body))
}

// DeleteService deletes a service from NagiosXI
func DeleteService(config Config, service Service) {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + "/" + config.BasePath + "/config/service?apikey=" +
		config.APIKey + "&pretty=1&host_name=" + service.HostName + "&service_description=" +
		url.QueryEscape(service.ServiceDescription))
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Deleting service %q for host %q:\n%s", service.ServiceDescription, service.HostName, string(body))
}
