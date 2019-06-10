package nagiosxi

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/gorilla/schema"
)

// Config holding NagiosXI configuration
type Config struct {
	APIKey   string `yaml:"apiKey"`
	BasePath string `yaml:"basePath"`
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Protocol string `yaml:"protocol"`
}

// ApplyConfig applies changes made previously. It also asks for user confirmation to make sure the user wants to do it.
func ApplyConfig(config Config) error {
	var choice string
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\n\nAre you sure you want to apply current NagiosXI configuration [y/N] ? ")
	choice, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	choice = strings.ToLower(strings.TrimSpace(choice))

	if choice == "y" || choice == "yes" {
		resp, err := http.Get(config.Protocol + "://" + config.Host + "/" + config.BasePath + "/system/applyconfig?apikey=" + config.APIKey)
		if err != nil {
			return fmt.Errorf("Error while making POST request to NagiosXI API: %s", err)
		}

		defer resp.Body.Close()

		fmt.Println("Configuration applied !")
	} else {
		fmt.Println("Not applying configuration.")
	}

	return nil
}

// BoolToStr converts boolean to string
func BoolToStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

// EncodeStringArray encodes string array in order to be use in NagiosXI API calls
func EncodeStringArray(v reflect.Value) string {
	var s string

	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			s += v.Index(i).Interface().(string) + ","
		}
	}

	if len(s) == 0 {
		return s
	}
	return s[:len(s)-1]

}

// EncodeStringArrayForDeletion encodes array in order to be compatible for a deletion request
func EncodeStringArrayForDeletion(array []string, tag string) string {
	var s string

	for _, elem := range array {
		s += "&" + tag + "[]=" + elem
	}

	if len(s) == 0 {
		return s
	}
	return s[1:]
}

// InitEncoder initializes an encoder and configure it
func InitEncoder() schema.Encoder {
	Encoder := schema.NewEncoder()
	Encoder.RegisterEncoder([]string{}, EncodeStringArray)
	return *Encoder
}

// IsExtensionValid indicates if the given filename has a valid extension
func IsExtensionValid(file string, validExtensions []string) bool {
	extension := file[strings.LastIndex(file, ".")+1:]
	for _, ext := range validExtensions {
		if extension == ext {
			return true
		}
	}
	return false
}
