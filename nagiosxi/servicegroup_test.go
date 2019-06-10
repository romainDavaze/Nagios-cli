package nagiosxi

import (
	"os"
	"reflect"
	"testing"
)

func TestAddServicegroup(t *testing.T) {
	sg := Servicegroup{
		Alias: "Test servicegroup",
		Name:  "testservicegroup",
	}

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	err := AddServicegroup(config, sg, true)
	if err != nil {
		t.Error(err)
	}

	servicegroup, err := GetServicegroup(config, sg.Name)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(sg, servicegroup) {
		t.Error("Mismatch between retrieved servicegroup and added servicegroup")
	}
}

func TestGetServicegroup(t *testing.T) {
	sg := "testservicegroup"

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	servicegroup, err := GetServicegroup(config, sg)

	if err != nil {
		t.Error(err)
	}

	if servicegroup.Name != sg {
		t.Error("Mismatch between requested servicegroup and NagiosXI's response")
	}
}

func TestDeleteServicegroup(t *testing.T) {
	sg := Servicegroup{
		Alias: "Test servicegroup",
		Name:  "testservicegroup",
	}

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	err := AddServicegroup(config, sg, true)
	if err != nil {
		t.Error(err)
	}
	err = DeleteServicegroup(config, sg)
	if err != nil {
		t.Error(err)
	}

	_, err = GetServicegroup(config, sg.Name)
	if err == nil {
		t.Error(err)
	}

}

func TestParseServicegroups(t *testing.T) {
	sg := []Servicegroup{
		Servicegroup{
			Alias:               "Servicegroup 1",
			Name:                "servicegroup1",
			ServicegroupMembers: nil,
		},
		Servicegroup{
			Alias:               "Servicegroup 2",
			Name:                "servicegroup2",
			ServicegroupMembers: []string{"servicegroup1"},
		},
	}

	dir, _ := os.Getwd()
	servicegroups, err := ParseServicegroups(dir + "/../test/testdata/servicegroups.yml")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(sg, servicegroups) {
		t.Error("Mismatch between parsed servicegroups and declared ones")
	}
}
