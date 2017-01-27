// Copyright © 2017 Drud Technology LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"log"

	prepo "github.com/drud/prepo/pkg"
	"github.com/spf13/cobra"
)

var cfgFile string

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "prepo",
	Short: "prep repos for lift off",
	Long:  `Prepo is a tool for setting up your github repositories with all the issues you could want.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		// make sure there is a prepo config file before proceeding
		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			log.Fatalf("File does not exist: %s", cfgFile)
		}
		// the only argument should be the repo in the form "org/repo"
		if len(args) != 1 {
			log.Fatal("incorrect number of arguments passed")
		}

	},
	Run: func(cmd *cobra.Command, args []string) {
		// get the contents of the prepo config file
		fileBytes, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			log.Fatalln("Could not read config file:", err)
		}
		// unmarshal the yaml into aPrepoConfig struct
		prepoConfig, err := prepo.GetPrepoConfig(fileBytes)
		if err != nil {
			log.Fatalln(err)
		}
		// use GITHUB_TOKEN env var to create an authed github client
		client, err := prepo.GetGithubClient()
		if err != nil {
			log.Fatal(err)
		}
		// loop through labels and add them to the target repo, ignores preexisting labels
		err = prepo.AddLabels(client, args[0], prepoConfig.Labels)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Repo updated successfully!")
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "prepo.yaml", "file that defines what will be changed in the target repo")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	return
}
