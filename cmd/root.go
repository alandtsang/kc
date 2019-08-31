/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/alandtsang/kc/manager"
	"github.com/alandtsang/kc/resources"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kc",
	Short: "kc controls the Kubernetes cluster manager",
	Long: `kc controls the Kubernetes cluster manager.

Find more information at: https://github.com/alandtsang/kc`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kc")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func NewKCManager(action resources.Action, namespace, resource, name string) {
	fmt.Printf("action: %v, namespace: %s, resource: %s, name: %s\n", action, namespace, resource, name)
	clusterName := "k8s-test2"
	var kcmanager manager.KCManager
	kcmanager.Init(action, clusterName, namespace, resource, name)
	kcmanager.Do()
}

func validate(args []string) {
	argsLen := len(args)
	if argsLen == 0 {
		log.Fatalln("[Error] Empty resource are not allowed")
	} else if argsLen > 2 {
		log.Fatalln("[Error] Too many parameters are not allowed")
	}

	if _, ok := resources.ResourcesLists[args[0]]; !ok {
		log.Fatalf("[Error] Resource name %s are not allowed\n", args[0])
	}
}

func do(action resources.Action, namespace string, args []string) {
	var resource, name string
	resource = args[0]
	if len(args) == 2 {
		name = args[1]
	}
	NewKCManager(action, namespace, resource, name)
}
