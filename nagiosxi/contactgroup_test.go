package nagiosxi

import (
	"os"
	"reflect"
	"testing"
)

func TestGetContactgroup(t *testing.T) {
	c := "admins"

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	contactgroup, err := GetContactgroup(config, c)

	if err != nil {
		t.Error(err)
	}

	if contactgroup.Name != c {
		t.Error("Mismatch between request contactgroup the response")
	}
}

func TestAddContactgroup(t *testing.T) {
	c := Contactgroup{
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

	AddContactgroup(config, c, true)

	contactgroup, err := GetContactgroup(config, c.Name)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(c, contactgroup) {
		t.Error("Mismatch between retrieved contactgroup and added contactgroup")
	}
}

func TestDeleteContactgroup(t *testing.T) {
	c := Contactgroup{
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

	AddContactgroup(config, c, true)
	DeleteContactgroup(config, c)

	_, err := GetContactgroup(config, c.Name)
	if err == nil {
		t.Error(err)
	}

}

func TestParseContactgroups(t *testing.T) {
	c := []Contactgroup{
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

	if !reflect.DeepEqual(c, contactgroups) {
		t.Error("Mismatch between parsed contactgroups and declared ones")
	}
}
