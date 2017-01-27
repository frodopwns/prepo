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

package pkg

import (
	"fmt"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// PrepoConfig represents the content in the prepo config file
type PrepoConfig struct {
	Labels []github.Label
}

// AddLabels loops through the given labels and adds them to the github repo
func AddLabels(client *github.Client, repo string, LabelsIn []github.Label) error {
	fmt.Println("adding labels to", repo)
	s := strings.Split(repo, "/")
	org, repoName := s[0], s[1]
	for _, labelIn := range LabelsIn {
		_, _, err := client.Issues.CreateLabel(org, repoName, &labelIn)
		if err != nil {
			if !strings.Contains(err.Error(), "already_exists") {
				return err
			}
		}
	}

	return nil
}

// GetGithubClient returns an authed github client
func GetGithubClient() (*github.Client, error) {
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		return nil, fmt.Errorf("No GITHUB_TOKEN env var set")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return github.NewClient(tc), nil
}

// GetPrepoConfig creates a new prepo config struct and unmarshalls the config content into it
func GetPrepoConfig(b []byte) (*PrepoConfig, error) {
	pc := &PrepoConfig{}
	err := yaml.Unmarshal(b, pc)
	if err != nil {
		return nil, err
	}
	return pc, nil
}
