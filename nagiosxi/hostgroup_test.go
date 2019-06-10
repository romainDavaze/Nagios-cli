package nagiosxi

import (
	"os"
	"reflect"
	"testing"
)

func TestGetHostgroup(t *testing.T) {
	hg := "linux-servers"

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	hostgroup, err := GetHostgroup(config, hg)

	if err != nil {
		t.Error(err)
	}

	if hostgroup.Name != hg {
		t.Error("Mismatch between requested hostgroup and NagiosXI's response")
	}
}

func TestAddHostgroup(t *testing.T) {
	hg := Hostgroup{
		Alias:            "Test Hostgroup",
		HostgroupMembers: []string{"linux-servers"},
		Members:          []string{"localhost"},
		Name:             "testhostgroup",
	}

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	err := AddHostgroup(config, hg, true)
	if err != nil {
		t.Error(err)
	}

	hostgroup, err := GetHostgroup(config, hg.Name)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(hg, hostgroup) {
		t.Error("Mismatch between retrieved hostgroup and added hostgroup")
	}
}

func TestDeleteHostgroup(t *testing.T) {
	hg := Hostgroup{
		Alias:            "Test Hostgroup",
		HostgroupMembers: []string{"linux-servers"},
		Members:          []string{"localhost"},
		Name:             "testhostgroup",
	}

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	err := AddHostgroup(config, hg, true)
	if err != nil {
		t.Error(err)
	}
	err = DeleteHostgroup(config, hg)
	if err != nil {
		t.Error(err)
	}

	_, err = GetHostgroup(config, hg.Name)
	if err == nil {
		t.Error(err)
	}

}

func TestParseHostgroups(t *testing.T) {
	hg := []Hostgroup{
		Hostgroup{
			Alias:            "Hostgroup 1",
			HostgroupMembers: []string{"linux-servers"},
			Members:          []string{"localhost"},
			Name:             "hostgroup1",
		},
		Hostgroup{
			Alias:            "Hostgroup 2",
			HostgroupMembers: nil,
			Members:          []string{"localhost"},
			Name:             "hostgroup2",
		},
	}

	dir, _ := os.Getwd()
	hostgroups, err := ParseHostgroups(dir + "/../test/testdata/hostgroups.yml")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(hg, hostgroups) {
		t.Error("Mismatch between parsed hostgroups and declared ones")
	}
}
