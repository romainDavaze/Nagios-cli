package nagiosxi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Contact represents the NagiosXI contact object
type Contact struct {
	Alias                       string   `schema:"alias" yaml:"alias"`
	ContactGroups               []string `schema:"contactgroups" yaml:"contactgroups"`
	Email                       string   `schema:"email" yaml:"email"`
	HostNotificationsEnabled    string   `schema:"host_notifications_enabled" yaml:"hostNotificationsEnabled"`
	HostNotificationsOptions    []string `schema:"host_notification_options" yaml:"hostNotificationsOptions"`
	HostNotificationCommands    []string `schema:"host_notification_commands" yaml:"hostNotificationCommands"`
	HostNotificationsPeriod     string   `schema:"host_notification_period" yaml:"hostNotificationsPeriod"`
	Members                     []string `schema:"members" yaml:"members"`
	Name                        string   `schema:"contact_name" yaml:"name"`
	ServiceNotificationsEnabled string   `schema:"service_notifications_enabled" yaml:"serviceNotificationsEnabled"`
	ServiceNotificationsOptions []string `schema:"service_notification_options" yaml:"serviceNotificationsOptions"`
	ServiceNotificationCommands []string `schema:"service_notification_commands" yaml:"serviceNotificationCommands"`
	ServiceNotificationsPeriod  string   `schema:"service_notification_period" yaml:"serviceNotificationsPeriod"`
	Templates                   []string `schema:"use" yaml:"templates"`
}

// AddContact adds a contact to NagiosXI
func AddContact(config Config, contact Contact) {
	values := make(map[string][]string)

	encoder := InitEncoder()
	err := encoder.Encode(contact, values)
	if err != nil {
		log.Fatalf("Error while encoding contact %q: %s", contact.Name, err)
	}

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+"/"+config.BasePath+"/config/contact?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		log.Fatalf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding contact %q:\n%s", contact.Name, string(body))
}

// DeleteContact deletes a contacts from NagiosXI
func DeleteContact(config Config, contact Contact) {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + "/" + config.BasePath + "/config/contact?apikey=" +
		config.APIKey + "&pretty=1&contact_name=" + contact.Name)
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Deleting contact %q:\n%s", contact.Name, string(body))
}

// ParseContacts parses NagiosXI contacts from a given yaml file
func ParseContacts(file string) []Contact {
	var objects map[string][]map[string]interface{}

	content, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(content, &objects)

	obj := objects["contacts"]
	if len(obj) == 0 {
		log.Fatal("There is no contacts object in the given file")
	}

	var contacts []Contact
	mapstructure.Decode(obj, &contacts)

	return contacts
}
