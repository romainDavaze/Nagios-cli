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

// Contactgroup represents the NagiosXI contactgroup object
type Contactgroup struct {
	Alias               string   `json:"alias" schema:"alias,omitempty" yaml:"alias"`
	Members             []string `json:"members" schema:"members,omitempty" yaml:"members"`
	ContactgroupMembers []string `json:"contactgroup_members" schema:"contactgroup_members,omitempty" yaml:"contactgroupMembers"`
	Name                string   `json:"contactgroup_name" schema:"contactgroup_name,omitempty" yaml:"name"`
}

// AddContactgroup adds a contactgroup to NagiosXI
func AddContactgroup(config Config, contactgroup Contactgroup, force bool) error {
	values := make(map[string][]string)

	encoder := InitEncoder()
	err := encoder.Encode(contactgroup, values)
	if err != nil {
		return fmt.Errorf("Error while encoding contactgroup %q: %s", contactgroup.Name, err)
	}

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+":"+strconv.Itoa(int(config.Port))+"/"+config.BasePath+"/config/contactgroup?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		return fmt.Errorf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding contactgroup %q:\n%s", contactgroup.Name, string(body))

	return nil
}

// DeleteContactgroup deletes a contactgroups from NagiosXI
func DeleteContactgroup(config Config, contactgroup Contactgroup) error {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/contactgroup?apikey=" +
		config.APIKey + "&pretty=1&contactgroup_name=" + contactgroup.Name)
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Deleting contactgroup %q:\n%s", contactgroup.Name, string(body))

	return nil
}

// GetContactgroup retrives a contactgroup from NagiosXI
func GetContactgroup(config Config, contactgroupName string) (Contactgroup, error) {
	contactgroups := []Contactgroup{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/contactgroup?apikey=" +
		config.APIKey + "&pretty=1&contactgroup_name=" + contactgroupName)
	resp, err := http.Get(fullURL)
	if err != nil {
		return Contactgroup{}, fmt.Errorf("Error while retrieving %s contactgroup from NagiosXI: %s", contactgroupName, err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &contactgroups)
	if err != nil {
		return Contactgroup{}, fmt.Errorf("Error while unmarshalling %s contactgroup from NagiosXI: %s", contactgroupName, err)
	}

	if len(contactgroups) == 0 {
		return Contactgroup{}, fmt.Errorf("Could not retrieve contactgroup %s from NagiosXI", contactgroupName)
	}

	return contactgroups[0], nil
}

// ParseContactgroups parses NagiosXI contactgroups from a given yaml file
func ParseContactgroups(file string) ([]Contactgroup, error) {
	contactgroups := []Contactgroup{}
	var objects map[string][]map[string]interface{}

	content, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(content, &objects)

	obj := objects["contactgroups"]
	if len(obj) == 0 {
		return contactgroups, fmt.Errorf("There is no contactgroups object in the given file")
	}

	mapstructure.Decode(obj, &contactgroups)

	return contactgroups, nil
}
