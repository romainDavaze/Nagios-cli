package nagiosxi

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestGetHost(t *testing.T) {
	h := "localhost"

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	host, err := GetHost(config, h)

	if err != nil {
		t.Error(err)
	}

	if host.Name != h {
		t.Error("Mismatch between requested host and NagiosXI's response")
	}
}

func TestAddHost(t *testing.T) {
	h := Host{
		Address:      "192.168.1.1",
		CheckCommand: "check_none",
		CheckPeriod:  "24x7",
		DisplayName:  "Test host",
		Name:         "testHost",
	}

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	err := AddHost(config, h, true)
	if err != nil {
		t.Error(err)
	}

	host, err := GetHost(config, h.Name)

	fmt.Println(h)
	fmt.Println(host)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(h, host) {
		t.Error("Mismatch between retrieved host and added host")
	}
}

// func TestDeleteHost(t *testing.T) {
// 	c := Host{
// 		Address:      "192.168.1.1",
// 		CheckCommand: "check_none",
// 		DisplayName:  "Test host",
// 		Name:         "testHost",
// 	}

// 	config := Config{
// 		APIKey:   os.Getenv("API_KEY"),
// 		BasePath: "nagiosxi/api/v1",
// 		Host:     "localhost",
// 		Port:     8080,
// 		Protocol: "http",
// 	}

// 	err := AddHost(config, c, true)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = DeleteHost(config, c)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	_, err = GetHost(config, c.Name)
// 	if err == nil {
// 		t.Error(err)
// 	}

// }

// func TestParseHosts(t *testing.T) {
// 	c := []Host{
// 		// Host{
// 		// 	Name:     "testHost",
// 		// 	HostLine: "ping test",
// 		// },
// 		// Host{
// 		// 	Name:     "testHost2",
// 		// 	HostLine: "ping test2",
// 		// },
// 	}

// 	dir, _ := os.Getwd()
// 	hosts, err := ParseHosts(dir + "/../test/testdata/hosts.yml")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if !reflect.DeepEqual(c, hosts) {
// 		t.Error("Mismatch between parsed hosts and declared ones")
// 	}
// }
