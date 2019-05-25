package cmd

import (
	"fmt"
	"nagios-cli/nagios"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const basePath = "api/v1/config"

var nagiosConfig nagios.Config

// App configuration file
var cfgFile string

// File containing Nagios objects to parse
var nagiosFile string

var rootCmd = &cobra.Command{
	Use:   "nagios-cli",
	Short: "CLI to interact with Nagios API",
	Long: `Command Line Interface (CLI) to interact with Nagios API.
	
	It allow to directly manage Nagios objects from the command line`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nagios-cli.yaml)")

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
		viper.SetConfigName(".nagios-cli")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	nagiosConfig = nagios.Config{
		APIKey:   viper.GetString("apiKey"),
		Host:     viper.GetString("nagiosHost"),
		BasePath: basePath,
	}
}
