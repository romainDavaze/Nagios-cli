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

// Servicegroup represents the NagiosXI servicegroup object
type Servicegroup struct {
	Alias               string   `json:"alias" schema:"alias,omitempty" yaml:"alias"`
	Members             []string `json:"members" schema:"members,omitempty" yaml:"members"`
	ServicegroupMembers []string `json:"servicegroup_members" schema:"servicegroup_members,omitempty" yaml:"servicegroupMembers"`
	Name                string   `json:"servicegroup_name" schema:"servicegroup_name,omitempty" yaml:"name"`
}

// Encode encodes a servicegroup into a map[string][]string
func (servicegroup *Servicegroup) Encode(force bool) (map[string][]string, error) {
	values := make(map[string][]string)

	encoder := InitEncoder()
	err := encoder.Encode(servicegroup, values)

	values["force"] = []string{BoolToStr(force)}

	return values, err
}

// AddServicegroup adds a servicegroup to NagiosXI
func AddServicegroup(config Config, servicegroup Servicegroup, force bool) error {
	values, err := servicegroup.Encode(force)
	if err != nil {
		return fmt.Errorf("Error while encoding contact %q: %s", servicegroup.Name, err)
	}

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+":"+strconv.Itoa(int(config.Port))+"/"+config.BasePath+"/config/servicegroup?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		return fmt.Errorf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding servicegroup %q:\n%s", servicegroup.Name, string(body))

	return nil
}

// DeleteServicegroup deletes a servicegroups from NagiosXI
func DeleteServicegroup(config Config, servicegroup Servicegroup) error {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/servicegroup?apikey=" +
		config.APIKey + "&pretty=1&servicegroup_name=" + servicegroup.Name)
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Deleting servicegroup %q:\n%s", servicegroup.Name, string(body))

	return nil
}

// GetServicegroup retrives a serviceGroup from NagiosXI
func GetServicegroup(config Config, serviceGroupName string) (Servicegroup, error) {
	serviceGroups := []Servicegroup{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + ":" + strconv.Itoa(int(config.Port)) + "/" + config.BasePath + "/config/servicegroup?apikey=" +
		config.APIKey + "&pretty=1&servicegroup_name=" + serviceGroupName)
	resp, err := http.Get(fullURL)
	if err != nil {
		return Servicegroup{}, fmt.Errorf("Error while retrieving %s serviceGroup from NagiosXI: %s", serviceGroupName, err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &serviceGroups)
	if err != nil {
		return Servicegroup{}, fmt.Errorf("Error while unmarshalling %s serviceGroup from NagiosXI: %s", serviceGroupName, err)
	}

	if len(serviceGroups) == 0 {
		return Servicegroup{}, fmt.Errorf("Could not retrieve serviceGroup %s from NagiosXI", serviceGroupName)
	}

	return serviceGroups[0], nil
}

// ParseServicegroups parses NagiosXI servicegroups from a given yaml file
func ParseServicegroups(file string) ([]Servicegroup, error) {
	var objects map[string][]map[string]interface{}
	servicegroups := []Servicegroup{}

	content, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(content, &objects)

	obj := objects["servicegroups"]
	if len(obj) == 0 {
		return servicegroups, fmt.Errorf("There is no servicegroups object in the given file")
	}

	mapstructure.Decode(obj, &servicegroups)

	return servicegroups, nil
}
