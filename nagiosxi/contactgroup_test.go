package nagiosxi

import (
	"os"
	"reflect"
	"testing"
)

func TestGetContactgroup(t *testing.T) {
	cg := "admins"

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	contactgroup, err := GetContactgroup(config, cg)

	if err != nil {
		t.Error(err)
	}

	if contactgroup.Name != cg {
		t.Error("Mismatch between requested contactgroup and NagiosXI's response")
	}
}

func TestAddContactgroup(t *testing.T) {
	cg := Contactgroup{
		Alias:               "Test contactgroup",
		Name:                "testcontactgroup",
		Members:             []string{"nagiosadmin"},
		ContactgroupMembers: []string{"admins", "xi_contactgroup_all"},
	}

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	err := AddContactgroup(config, cg, true)
	if err != nil {
		t.Error(err)
	}

	contactgroup, err := GetContactgroup(config, cg.Name)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(cg, contactgroup) {
		t.Error("Mismatch between retrieved contactgroup and added contactgroup")
	}
}

func TestDeleteContactgroup(t *testing.T) {
	cg := Contactgroup{
		Alias:               "Test contactgroup",
		Name:                "testcontactgroup",
		Members:             []string{"nagiosadmin"},
		ContactgroupMembers: []string{"admins", "xi_contactgroup_all"},
	}

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	err := AddContactgroup(config, cg, true)
	if err != nil {
		t.Error(err)
	}
	err = DeleteContactgroup(config, cg)
	if err != nil {
		t.Error(err)
	}

	_, err = GetContactgroup(config, cg.Name)
	if err == nil {
		t.Error(err)
	}

}

func TestParseContactgroups(t *testing.T) {
	cg := []Contactgroup{
		Contactgroup{
			Alias:               "Contactgroup 1",
			Name:                "contactgroup1",
			Members:             []string{"nagiosadmin"},
			ContactgroupMembers: []string{"xi_contactgroup_all", "admins"},
		},
		Contactgroup{
			Alias:               "Contactgroup 2",
			Name:                "contactgroup2",
			Members:             []string{"nagiosadmin"},
			ContactgroupMembers: nil,
		},
	}

	dir, _ := os.Getwd()
	contactgroups, err := ParseContactgroups(dir + "/../test/testdata/contactgroups.yml")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(cg, contactgroups) {
		t.Error("Mismatch between parsed contactgroups and declared ones")
	}
}
