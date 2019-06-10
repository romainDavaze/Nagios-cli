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

// Hostgroup represents the NagiosXI hostgroup object
type Hostgroup struct {
	Alias            string   `json:"alias" schema:"alias,omitempty" yaml:"alias"`
	Members          []string `json:"members" schema:"members,omitempty" yaml:"members"`
	HostgroupMembers []string `json:"hostgroup_members" schema:"hostgroup_members,omitempty" yaml:"hostgroupMembers"`
	Name             string   `json:"hostgroup_name" schema:"hostgroup_name,omitempty" yaml:"name"`
}

// Encode encodes a hostgroup into a map[string][]string
func (hostgroup *Hostgroup) Encode(force bool) (map[string][]string, error) {
	values := make(map[string][]string)

	encoder := InitEncoder()
	err := encoder.Encode(hostgroup, values)

	values["force"] = []string{BoolToStr(force)}

	return values, err
}

// AddHostgroup adds a hostgroup to NagiosXI
func AddHostgroup(config Config, hostgroup Hostgroup, force bool) error {
	values, err := hostgroup.Encode(force)
	if err != nil {
		return fmt.Errorf("Error while encoding hostgroup %q: %s", hostgroup.Name, err)
	}

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+":"+strconv.Itoa(int(config.Port))+"/"+config.BasePath+"/config/hostgroup?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		return fmt.Errorf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding hostgroup %q:\n%s", hostgroup.Name, string(body))

	return nil
}

// DeleteHostgroup deletes a hostgroups from NagiosXI
func DeleteHostgroup(config Config, hostgroup Hostgroup) error {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/hostgroup?apikey=" +
		config.APIKey + "&pretty=1&hostgroup_name=" + hostgroup.Name)
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Deleting hostgroup %q:\n%s", hostgroup.Name, string(body))

	return nil
}

// GetHostgroup retrives a hostgroup from NagiosXI
func GetHostgroup(config Config, HostgroupName string) (Hostgroup, error) {
	hostgroups := []Hostgroup{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/hostgroup?apikey=" +
		config.APIKey + "&pretty=1&hostgroup_name=" + HostgroupName)
	resp, err := http.Get(fullURL)
	if err != nil {
		return Hostgroup{}, fmt.Errorf("Error while retrieving %s hostgroup from NagiosXI: %s", HostgroupName, err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &hostgroups)
	if err != nil {
		return Hostgroup{}, fmt.Errorf("Error while unmarshalling %s hostgroup from NagiosXI: %s", HostgroupName, err)
	}

	if len(hostgroups) == 0 {
		return Hostgroup{}, fmt.Errorf("Could not retrieve hostgroup %s from NagiosXI", HostgroupName)
	}

	return hostgroups[0], nil
}

// ParseHostgroups parses NagiosXI hostgroups from a given yaml file
func ParseHostgroups(file string) ([]Hostgroup, error) {
	hostgroups := []Hostgroup{}
	var objects map[string][]map[string]interface{}

	content, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(content, &objects)

	obj := objects["hostgroups"]
	if len(obj) == 0 {
		return hostgroups, fmt.Errorf("There is no hostgroups object in the given file")
	}

	mapstructure.Decode(obj, &hostgroups)

	return hostgroups, nil
}
