package nagiosxi

import (
	"os"
	"reflect"
	"testing"
)

func TestGetService(t *testing.T) {
	c := "localhost"
	s := "PING"

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	service, err := GetService(config, c, s)

	if err != nil {
		t.Error(err)
	}

	if service.ServiceDescription != s {
		t.Error("Mismatch between requested service and NagiosXI's response")
	}
}

func TestAddService(t *testing.T) {
	s := Service{
		CheckCommand:       "check_none",
		ConfigName:         "localhost",
		HostGroups:         []string{"linux-servers"},
		Hosts:              []string{"localhost"},
		ServiceDescription: "This is a check_none service",
	}

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	err := AddService(config, s, true)
	if err != nil {
		t.Error(err)
	}

	service, err := GetService(config, s.Hosts[0], s.ServiceDescription)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(s, service) {
		t.Error("Mismatch between retrieved service and added service")
	}
}

func TestDeleteService(t *testing.T) {
	s := Service{
		CheckCommand:       "check_none",
		ConfigName:         "localhost",
		HostGroups:         []string{"linux-servers"},
		Hosts:              []string{"localhost"},
		ServiceDescription: "This is a check_none service",
	}

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	err := AddService(config, s, true)
	if err != nil {
		t.Error(err)
	}
	err = DeleteService(config, s)
	if err != nil {
		t.Error(err)
	}

	_, err = GetService(config, s.Hosts[0], s.ServiceDescription)
	if err == nil {
		t.Error(err)
	}

}

func TestParseServices(t *testing.T) {
	s := []Service{
		Service{
			CheckCommand:         "check_ping",
			CheckCommandArgs:     []string{"3000,80%", "5000,100%"},
			CheckInterval:        "5",
			CheckPeriod:          "24x7",
			ConfigName:           "localhost",
			Contacts:             []string{"nagiosadmin"},
			ContactGroups:        []string{"admins", "xi_contactgroup_all"},
			DisplayName:          "Ping Service",
			HostGroups:           []string{"linux-servers"},
			Hosts:                []string{"localhost"},
			MaxCheckAttempts:     "2",
			NotificationInterval: "5",
			NotificationPeriod:   "24x7",
			RetryInterval:        "5",
			ServiceDescription:   "This is a ping service",
			Templates:            []string{"local-service", "generic-service"},
		},
		Service{
			CheckCommand:         "check_none",
			CheckCommandArgs:     nil,
			CheckInterval:        "",
			CheckPeriod:          "",
			ConfigName:           "",
			Contacts:             nil,
			ContactGroups:        nil,
			DisplayName:          "",
			HostGroups:           []string{"windows-servers"},
			Hosts:                []string{"localhost"},
			MaxCheckAttempts:     "",
			NotificationInterval: "",
			NotificationPeriod:   "",
			RetryInterval:        "",
			ServiceDescription:   "This is a check_none service",
			Templates:            nil,
		},
	}

	dir, _ := os.Getwd()
	services, err := ParseServices(dir + "/../test/testdata/services.yml")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(s, services) {
		t.Error("Mismatch between parsed services and declared ones")
	}
}
