package nagiosxi

import (
	"os"
	"reflect"
	"testing"
)

func TestGetContact(t *testing.T) {
	c := "nagiosadmin"

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	contact, err := GetContact(config, c)

	if err != nil {
		t.Error(err)
	}

	if contact.Name != c {
		t.Error("Mismatch between request contact the response")
	}
}

func TestAddContact(t *testing.T) {
	c := Contact{
		Alias:                       "Test Contact",
		HostNotificationCommands:    []string{"notify-host-by-email"},
		HostNotificationOptions:     "f,s",
		HostNotificationPeriod:      "24x7",
		HostNotificationsEnabled:    "0",
		Name:                        "testcontact",
		ServiceNotificationCommands: []string{"notify-host-by-email"},
		ServiceNotificationOptions:  "r,n",
		ServiceNotificationPeriod:   "24x7",
		ServiceNotificationsEnabled: "1",
	}

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	AddContact(config, c, true)

	contact, err := GetContact(config, c.Name)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(c, contact) {
		t.Error("Mismatch between retrieved contact and added contact")
	}
}

func TestDeleteContact(t *testing.T) {
	c := Contact{
		Alias:                       "Test Contact",
		HostNotificationCommands:    []string{"notify-host-by-email"},
		HostNotificationOptions:     "f,s",
		HostNotificationPeriod:      "24x7",
		HostNotificationsEnabled:    "0",
		Name:                        "testcontact",
		ServiceNotificationCommands: []string{"notify-host-by-email"},
		ServiceNotificationOptions:  "r,n",
		ServiceNotificationPeriod:   "24x7",
		ServiceNotificationsEnabled: "1",
	}

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	AddContact(config, c, true)
	DeleteContact(config, c)

	_, err := GetContact(config, c.Name)
	if err == nil {
		t.Error(err)
	}

}

func TestParseContacts(t *testing.T) {
	c := []Contact{
		Contact{
			Alias:                       "Contact 1",
			ContactGroups:               []string{"xi_contactgroup_all", "admins"},
			Email:                       "root@localhost",
			HostNotificationCommands:    []string{"notify-host-by-email"},
			HostNotificationOptions:     "d,u",
			HostNotificationPeriod:      "24x7",
			HostNotificationsEnabled:    "1",
			Members:                     nil,
			Name:                        "contact1",
			ServiceNotificationCommands: []string{"notify-host-by-email"},
			ServiceNotificationOptions:  "c,w",
			ServiceNotificationPeriod:   "24x7",
			ServiceNotificationsEnabled: "0",
			Templates:                   nil,
		},
		Contact{
			Alias:                       "Contact 2",
			Email:                       "",
			HostNotificationCommands:    []string{"notify-host-by-email"},
			HostNotificationOptions:     "f,s",
			HostNotificationPeriod:      "24x7",
			HostNotificationsEnabled:    "0",
			Members:                     nil,
			Name:                        "contact2",
			ServiceNotificationCommands: []string{"notify-host-by-email"},
			ServiceNotificationOptions:  "r,s",
			ServiceNotificationPeriod:   "24x7",
			ServiceNotificationsEnabled: "1",
			Templates:                   nil,
		},
	}

	dir, _ := os.Getwd()
	contacts, err := ParseContacts(dir + "/../test/testdata/contacts.yml")

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(contacts, c) {
		t.Error("Mismatch between parsed contacts and declared ones")
	}
}
