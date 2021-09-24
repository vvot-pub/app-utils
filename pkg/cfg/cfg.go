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
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

// ConfigError represents all configuration related errors
type ConfigError struct {
	Message string
	Cause error
}

func (err ConfigError) Error() string {
	return fmt.Sprintf("%s: %v", err.Message, err.Cause)
}

// LoadConfig loads configuration from the given file into the specified output structure.
func LoadConfig(configFile string, outConf interface{}) (error) {
	configData, err := ReadConfigFile(configFile)
	if err != nil {
		return err
	}
	s :=  string(configData)
	_ = s
	return ReadConfig(configData, outConf)
}

// LoadAppConfig application configuration structure using the specified application name.
// The application configuration file assumed to be in $HOME/.config/<appName>.yaml file.
func LoadAppConfig(appName string, outConf interface{}) (error) {
	configData, err := ReadAppConfigFile(appName)
	if err != nil {
		return err
	}
	return ReadConfig(configData, outConf)
}

// ReadConfigFile retrieves configuration data using the specified file name.
func ReadConfigFile(configFile string) ([]byte, error) {
	result, err := ioutil.ReadFile(configFile)
	if err != nil {
		err := ConfigError{fmt.Sprintf("Cannot cannot read configuration file %s", configFile), err}
		return make([]byte, 0), err
	}
	return result, nil
}

// ReadAppConfigFile retrieves configuration data using the specified application name.
// The application configuration file assumed to be in $HOME/.config/<appName>.yaml file.
func ReadAppConfigFile(appName string) ([]byte, error) {
	currentUser, err := user.Current()
	if err != nil {
		err := ConfigError{fmt.Sprintf("Cannot retrive current user for appplication %s", appName), err}
		return make([]byte, 0), err
	}

	homeDir := currentUser.HomeDir
	configFile := filepath.Join(homeDir, ".config", appName+".yaml")
	return ReadConfigFile(configFile)
}

// ReadConfig reads configuration from the given buffer into the specified output structure.
func ReadConfig(buffer []byte, outConf interface{}) (error) {
	err := yaml.Unmarshal(buffer, outConf)
	if err != nil {
		return &ConfigError{"Cannot read configuration", err}
	}
	return nil
}
