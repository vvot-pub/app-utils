/*
Copyright 2021 Vladimir Votiakov.

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

package cfg

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

type TaskInput struct {
	Files  []string
	Source string
}

type TaskOutput struct {
	Name       string   `yaml:"name"`
	Id         int      `yaml:"id"`
	Parameters []string `yaml:"parameters"`
}

type TaskConfig struct {
	Input  TaskInput  `yaml:"taskInput"`
	Output TaskOutput `yaml:"taskOutput"`
}

var confStruct = `
taskInput:
  files: 
    - '*.xml'
    - '*.lib'
  source: /test/data
taskOutput:
  name: aggregator
  id: 100
  parameters:
    - one
    - two
    - three
`

func TestConfigFromString(t *testing.T) {
	config := TaskConfig{}
	err := ReadConfig([]byte(confStruct), &config)
	if err != nil {
		t.Fatalf("Cannot read config: %v", err)
	}

	refConfig := TaskConfig{
		 Input: TaskInput{
			 Files: []string{"*.xml", "*.lib"},
			 Source: "/test/data",
		 },
		 Output: TaskOutput{
			 Name: "aggregator",
			 Id: 100,
			 Parameters: []string{"one", "two", "three"},
		 },
	}
	if !reflect.DeepEqual(config, refConfig) {
		t.Error("Loaded and reference content are not the same")
	}
}

func TestConfigFromFile(t *testing.T) {
	resourcesDir := getTestResourcesDir()
	if resourcesDir == "" {
		t.Fatal("Cannot retrieve test directory")
	}
	configFile := filepath.Join(resourcesDir, "taskConfig.yaml")
	config := TaskConfig{}
	err := LoadConfig(configFile, &config)
	if err != nil {
		t.Errorf("Cannot load test config file: %v", err)
		return
	}
	refConfig := TaskConfig{
		Input: TaskInput{
			Files: []string{"*.json", "*.txt", "*.doc"},
			Source: "/test/data",
		},
		Output: TaskOutput{
			Name: "copier",
			Id: 1000,
			Parameters: []string{"aaa", "bbb" , "ccc"},
		},
	}
	if !reflect.DeepEqual(config, refConfig) {
		t.Error("Loaded and reference content are not the same")
	}
}

func getTestResourcesDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	testDir := filepath.Dir(filename)
	return filepath.Join(testDir, "../../test/resources/pkg/cfg")
}