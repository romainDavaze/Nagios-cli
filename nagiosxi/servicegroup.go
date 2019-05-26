package nagiosxi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Servicegroup represents the NagiosXI servicegroup object
type Servicegroup struct {
	Alias               string   `schema:"alias" yaml:"alias"`
	Members             []string `schema:"members" yaml:"members"`
	ServicegroupMembers []string `schema:"servicegroup_members" yaml:"servicegroupMembers"`
	Name                string   `schema:"servicegroup_name" yaml:"name"`
}

// AddServicegroup adds a servicegroup to NagiosXI
func AddServicegroup(config Config, servicegroup Servicegroup) {
	values := make(map[string][]string)

	encoder := InitEncoder()
	err := encoder.Encode(servicegroup, values)
	if err != nil {
		log.Fatalf("Error while encoding servicegroup %q: %s", servicegroup.Name, err)
	}

	resp, err := http.PostForm(config.Protocol+"://"+config.Host+"/"+config.BasePath+"/config/servicegroup?apikey="+config.APIKey+"&pretty=1", values)
	if err != nil {
		log.Fatalf("Error while making POST request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Adding servicegroup %q:\n%s", servicegroup.Name, string(body))
}

// DeleteServicegroup deletes a servicegroups from NagiosXI
func DeleteServicegroup(config Config, servicegroup Servicegroup) {
	client := &http.Client{}

	fullURL := fmt.Sprintf(config.Protocol + "://" + config.Host + "/" + config.BasePath + "/config/servicegroup?apikey=" +
		config.APIKey + "&pretty=1&servicegroup_name=" + servicegroup.Name)
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error while making DELETE request to NagiosXI API: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Deleting servicegroup %q:\n%s", servicegroup.Name, string(body))
}

// ParseServicegroups parses NagiosXI servicegroups from a given yaml file
func ParseServicegroups(file string) []Servicegroup {
	var objects map[string][]map[string]interface{}

	content, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(content, &objects)

	obj := objects["servicegroups"]
	if len(obj) == 0 {
		log.Fatal("There is no servicegroups object in the given file")
	}

	var servicegroups []Servicegroup
	mapstructure.Decode(obj, &servicegroups)

	return servicegroups
}
