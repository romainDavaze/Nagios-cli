package nagiosxi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Contactgroup represents the NagiosXI contactgroup object
type Contactgroup struct {
	Alias               string   `schema:"alias,omitempty" yaml:"alias"`
	Members             []string `schema:"members,omitempty" yaml:"members"`
	ContactgroupMembers []string `schema:"contactgroup_members,omitempty" yaml:"contactgroupMembers"`
	Name                string   `schema:"contactgroup_name,omitempty" yaml:"name"`
}

// AddContactgroup adds a contactgroup to NagiosXI
func AddContactgroup(config Config, contactgroup Contactgroup, force bool) {
	values := make(map[string][]string)

	encoder := InitEncoder()
	err := encoder.Encode(contactgroup, values)
	if err != nil {
		log.Fatalf("Error while encoding contactgroup %q: %s", contactgroup.Name, err)
	}

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+"/"+config.BasePath+"/config/contactgroup?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		log.Fatalf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding contactgroup %q:\n%s", contactgroup.Name, string(body))
}

// DeleteContactgroup deletes a contactgroups from NagiosXI
func DeleteContactgroup(config Config, contactgroup Contactgroup) {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + "/" + config.BasePath + "/config/contactgroup?apikey=" +
		config.APIKey + "&pretty=1&contactgroup_name=" + contactgroup.Name)
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Deleting contactgroup %q:\n%s", contactgroup.Name, string(body))
}

// ParseContactgroups parses NagiosXI contactgroups from a given yaml file
func ParseContactgroups(file string) []Contactgroup {
	var objects map[string][]map[string]interface{}

	content, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(content, &objects)

	obj := objects["contactgroups"]
	if len(obj) == 0 {
		log.Fatal("There is no contactgroups object in the given file")
	}

	var contactgroups []Contactgroup
	mapstructure.Decode(obj, &contactgroups)

	return contactgroups
}
