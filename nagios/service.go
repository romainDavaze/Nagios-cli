package nagios

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Service represents the Nagios service object
type Service struct {
	CheckCommand         string   `json:"check_command"`
	CheckCommandArgs     []string `json:"-"`
	CheckInterval        uint16   `json:"check_interval"`
	CheckPeriod          string   `json:"check_period"`
	Contacts             []string `json:"contacts"`
	ContactGroups        []string `json:"contact_groups"`
	HostName             string   `json:"host_name"`
	MaxCheckAttempts     uint16   `json:"max_check_attempts"`
	NotificationInterval uint16   `json:"notification_interval"`
	NotificationPeriod   string   `json:"notification_period"`
	RetryInterval        uint16   `json:"retry_interval"`
	ServiceDescription   string   `json:"service_description"`
}

// MarshalJSON customizes Service object json representation
func (s *Service) MarshalJSON() ([]byte, error) {
	type Alias Service

	var argsString string
	for _, arg := range s.CheckCommandArgs {
		argsString += "\\!" + arg
	}

	return json.Marshal(&struct {
		CheckCommand string `json:"check_command"`
		*Alias
	}{
		CheckCommand: s.CheckCommand + argsString,
		Alias:        (*Alias)(s),
	})
}

// AddService adds a service to Nagios
func AddService(config Config, service Service) {
	requestBody, _ := service.MarshalJSON()
	requestBody, _ = AddApplyConfigToJSON(requestBody)

	resp, err := http.Post("http://"+config.Host+"/"+config.BasePath+"?apikey="+config.APIKey, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error while making POST request to Nagios API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("Added service %s for host %s:\n %s", service.ServiceDescription, service.HostName, string(body))
}

// DeleteService deletes a service from Nagios
func DeleteService(config Config, service Service) {
	requestBody := []byte(`{"host_name": "` + service.HostName + `", "service_description": "` + service.ServiceDescription + `"}`)
	requestBody, _ = AddApplyConfigToJSON(requestBody)

	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", "http://"+config.Host+"/"+config.BasePath+"?apikey="+config.APIKey, bytes.NewBuffer(requestBody))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error while making DELETE request to Nagios API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Printf("Deleted service %q for host %q: %s", service.ServiceDescription, service.HostName, string(body))
}
