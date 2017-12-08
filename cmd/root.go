package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configurationFile string

var rootCommand = &cobra.Command{
	Use:   "kuzmich-runner",
	Short: "A brief description of your application",
	Long:  `TODO: info about kuzmich-runner`,
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCommand.PersistentFlags().StringVar(
		&configurationFile,
		"config",
		"",
		"config file (default is $HOME/.kuzmich-runner.yaml)",
	)
}

func initConfig() {
	if configurationFile != "" {
		viper.SetConfigFile(configurationFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".kuzmich-runner")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
