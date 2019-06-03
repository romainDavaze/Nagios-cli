package nagiosxi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Hostgroup represents the NagiosXI hostgroup object
type Hostgroup struct {
	Alias            string   `schema:"alias,omitempty" yaml:"alias"`
	Members          []string `schema:"members,omitempty" yaml:"members"`
	HostgroupMembers []string `schema:"hostgroup_members,omitempty" yaml:"hostgroupMembers"`
	Name             string   `schema:"hostgroup_name,omitempty" yaml:"name"`
}

// AddHostgroup adds a hostgroup to NagiosXI
func AddHostgroup(config Config, hostgroup Hostgroup, force bool) {
	values := make(map[string][]string)

	encoder := InitEncoder()
	err := encoder.Encode(hostgroup, values)
	if err != nil {
		log.Fatalf("Error while encoding hostgroup %q: %s", hostgroup.Name, err)
	}

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+"/"+config.BasePath+"/config/hostgroup?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		log.Fatalf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding hostgroup %q:\n%s", hostgroup.Name, string(body))
}

// DeleteHostgroup deletes a hostgroups from NagiosXI
func DeleteHostgroup(config Config, hostgroup Hostgroup) {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + "/" + config.BasePath + "/config/hostgroup?apikey=" +
		config.APIKey + "&pretty=1&hostgroup_name=" + hostgroup.Name)
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Deleting hostgroup %q:\n%s", hostgroup.Name, string(body))
}

// ParseHostgroups parses NagiosXI hostgroups from a given yaml file
func ParseHostgroups(file string) []Hostgroup {
	var objects map[string][]map[string]interface{}

	content, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(content, &objects)

	obj := objects["hostgroups"]
	if len(obj) == 0 {
		log.Fatal("There is no hostgroups object in the given file")
	}

	var hostgroups []Hostgroup
	mapstructure.Decode(obj, &hostgroups)

	return hostgroups
}
