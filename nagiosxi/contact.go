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

// Contact represents the NagiosXI contact object
type Contact struct {
	Alias                       string   `json:"alias" schema:"alias,omitempty" yaml:"alias"`
	ContactGroups               []string `json:"contactgroups" schema:"contactgroups,omitempty" yaml:"contactgroups"`
	Email                       string   `json:"email" schema:"email,omitempty" yaml:"email"`
	HostNotificationOptions     string   `json:"host_notification_options" schema:"host_notification_options,omitempty" yaml:"hostNotificationOptions"`
	HostNotificationCommands    []string `json:"host_notification_commands" schema:"host_notification_commands,omitempty" yaml:"hostNotificationCommands"`
	HostNotificationPeriod      string   `json:"host_notification_period" schema:"host_notification_period,omitempty" yaml:"hostNotificationPeriod"`
	HostNotificationsEnabled    string   `json:"host_notifications_enabled" schema:"host_notifications_enabled,omitempty" yaml:"hostNotificationsEnabled"`
	Members                     []string `json:"members" schema:"members,omitempty" yaml:"members"`
	Name                        string   `json:"contact_name" schema:"contact_name,omitempty" yaml:"name"`
	ServiceNotificationOptions  string   `json:"service_notification_options" schema:"service_notification_options,omitempty" yaml:"serviceNotificationOptions"`
	ServiceNotificationCommands []string `json:"service_notification_commands" schema:"service_notification_commands,omitempty" yaml:"serviceNotificationCommands"`
	ServiceNotificationPeriod   string   `json:"service_notification_period" schema:"service_notification_period,omitempty" yaml:"serviceNotificationPeriod"`
	ServiceNotificationsEnabled string   `json:"service_notifications_enabled" schema:"service_notifications_enabled,omitempty" yaml:"serviceNotificationsEnabled"`
	Templates                   []string `json:"use" schema:"use,omitempty" yaml:"templates"`
}

// AddContact adds a contact to NagiosXI
func AddContact(config Config, contact Contact, force bool) error {
	values := make(map[string][]string)

	encoder := InitEncoder()
	err := encoder.Encode(contact, values)
	if err != nil {
		return fmt.Errorf("Error while encoding contact %q: %s", contact.Name, err)
	}

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+":"+strconv.Itoa(int(config.Port))+"/"+config.BasePath+"/config/contact?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		return fmt.Errorf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding contact %q:\n%s", contact.Name, string(body))

	return nil
}

// DeleteContact deletes a contacts from NagiosXI
func DeleteContact(config Config, contact Contact) error {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/contact?apikey=" +
		config.APIKey + "&pretty=1&contact_name=" + contact.Name)
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Deleting contact %q:\n%s", contact.Name, string(body))

	return nil
}

// GetContact retrives a contact from NagiosXI
func GetContact(config Config, contactName string) (Contact, error) {
	contacts := []Contact{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/contact?apikey=" +
		config.APIKey + "&pretty=1&contact_name=" + contactName)
	resp, err := http.Get(fullURL)
	if err != nil {
		return Contact{}, fmt.Errorf("Error while retrieving %s contact from NagiosXI: %s", contactName, err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &contacts)
	if err != nil {
		return Contact{}, fmt.Errorf("Error while unmarshalling %s contact from NagiosXI: %s", contactName, err)
	}

	if len(contacts) == 0 {
		return Contact{}, fmt.Errorf("Could not retrieve contact %s from NagiosXI", contactName)
	}

	return contacts[0], nil
}

// ParseContacts parses NagiosXI contacts from a given yaml file
func ParseContacts(file string) ([]Contact, error) {
	contacts := []Contact{}
	var objects map[string][]map[string]interface{}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return contacts, fmt.Errorf("Error while reading file: %s", err)
	}
	yaml.Unmarshal(content, &objects)

	obj := objects["contacts"]
	if len(obj) == 0 {
		return contacts, fmt.Errorf("There is no contacts object in the given file")
	}

	mapstructure.Decode(obj, &contacts)

	return contacts, nil
}
