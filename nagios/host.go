package nagios

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Host represents the Nagios host object
type Host struct {
	Address              string   `json:"address"`
	CheckCommand         string   `json:"check_command"`
	CheckCommandArgs     []string `json:"-"`
	CheckPeriod          string   `json:"check_period"`
	Contacts             []string `json:"contacts"`
	ContactGroups        []string `json:"contact_groups"`
	HostName             string   `json:"host_name"`
	MaxCheckAttempts     uint16   `json:"max_check_attempts"`
	NotificationInterval uint16   `json:"notification_interval"`
	NotificationPeriod   string   `json:"notification_period"`
}

// MarshalJSON customizes Host object json representation
func (h *Host) MarshalJSON() ([]byte, error) {
	type Alias Host

	var argsString string
	for _, arg := range h.CheckCommandArgs {
		argsString += "\\!" + arg
	}

	return json.Marshal(&struct {
		CheckCommand string `json:"check_command"`
		*Alias
	}{
		CheckCommand: h.CheckCommand + argsString,
		Alias:        (*Alias)(h),
	})
}

// AddHost adds a host to Nagios
func AddHost(config Config, host Host) {
	requestBody, _ := host.MarshalJSON()
	requestBody, _ = AddApplyConfigToJSON(requestBody)

	resp, err := http.Post("http://"+config.Host+"/"+config.BasePath+"/host?apikey="+config.APIKey, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error while making POST request to Nagios API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("Added host %q: %s", host.HostName, string(body))
}

// DeleteHost deletes a hosts from Nagios
func DeleteHost(config Config, host Host) {
	requestBody := []byte(`{"host_name": "` + host.HostName + `"}`)
	requestBody, _ = AddApplyConfigToJSON(requestBody)

	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", "http://"+config.Host+"/"+config.BasePath+"/host?apikey="+config.APIKey, bytes.NewBuffer(requestBody))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error while making DELETE request to Nagios API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Printf("Deleted host %q: %s", host.HostName, string(body))
}
