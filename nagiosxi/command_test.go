package nagiosxi

import (
	"os"
	"reflect"
	"testing"
)

func TestGetCommand(t *testing.T) {
	c := "check_none"

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	command, err := GetCommand(config, c)

	if err != nil {
		t.Error(err)
	}

	if command.Name != c {
		t.Error("Mismatch between requested command and NagiosXI's response")
	}
}

func TestAddCommand(t *testing.T) {
	c := Command{
		Name:        "test",
		CommandLine: "check test",
	}

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	err := AddCommand(config, c, true)
	if err != nil {
		t.Error(err)
	}

	command, err := GetCommand(config, c.Name)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(c, command) {
		t.Error("Mismatch between retrieved command and added command")
	}
}

func TestDeleteCommand(t *testing.T) {
	c := Command{
		Name:        "test",
		CommandLine: "check test",
	}

	config := Config{
		APIKey:   os.Getenv("API_KEY"),
		BasePath: "nagiosxi/api/v1",
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	err := AddCommand(config, c, true)
	if err != nil {
		t.Error(err)
	}
	err = DeleteCommand(config, c)
	if err != nil {
		t.Error(err)
	}

	_, err = GetCommand(config, c.Name)
	if err == nil {
		t.Error(err)
	}

}

func TestParseCommands(t *testing.T) {
	c := []Command{
		Command{
			Name:        "testCommand",
			CommandLine: "ping test",
		},
		Command{
			Name:        "testCommand2",
			CommandLine: "ping test2",
		},
	}

	dir, _ := os.Getwd()
	commands, err := ParseCommands(dir + "/../test/testdata/commands.yml")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(c, commands) {
		t.Error("Mismatch between parsed commands and declared ones")
	}
}
