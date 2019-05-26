package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/romainDavaze/nagiosxi-cli/nagiosxi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const basePath = "nagiosxi/api/v1"

var applyConfig bool
var nagiosxiConfig nagiosxi.Config

// App configuration file
var cfgFile string

// File containing NagiosXI objects to parse
var objectsFile string

var rootCmd = &cobra.Command{
	Use:   "nagiosxi-cli",
	Short: "CLI to interact with NagiosXI API",
	Long:  `Command Line Interface (CLI) to interact with NagiosXI API`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
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

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s \n\n", viper.ConfigFileUsed())
	}

	nagiosxiConfig = nagiosxi.Config{
		APIKey:   viper.GetString("apiKey"),
		Host:     viper.GetString("nagiosxiHost"),
		BasePath: basePath,
		Protocol: viper.GetString("protocol"),
	}
}
