package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
)

const basePath = "nagiosxi/api/v1"

var applyConfig bool
var cfgFile string
var force bool
var nagiosxiConfig nagiosxi.Config
var objectsFile string
var validExtensions = []string{"yaml", "yml"}

var rootCmd = &cobra.Command{
	Use:   "nagiosxi-cli",
	Short: "CLI to interact with NagiosXI API",
	Long:  `Command Line Interface (CLI) to interact with NagiosXI API`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// GenMarkdownDocs generates markdown docs
func GenMarkdownDocs() {
	err := doc.GenMarkdownTree(rootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nagiosxi-cli.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".nagiosxi-cli")
	}

	viper.AutomaticEnv()

	viper.ReadInConfig()

	nagiosxiConfig = nagiosxi.Config{
		APIKey:   viper.GetString("apiKey"),
		Host:     viper.GetString("nagiosxiHost"),
		BasePath: basePath,
		Protocol: viper.GetString("protocol"),
	}
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("Requires a file containing NagiosXI objects")
	}
	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		return errors.New("Provided file does not exist")
	}
	if !nagiosxi.IsExtensionValid(args[0], validExtensions) {
		return fmt.Errorf("File format is not supported. Must be one of [ %s ]", strings.Join(validExtensions, " | "))
	}
	return nil
}
