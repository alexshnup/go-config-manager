/*
Package yaml provides configuration loading from yaml file.

Load(config interface{}) gets struct.

// Config struct
type Config struct {
	Debug   bool   `yaml:"debug"`
	Timeout int    `yaml:"timeout"`
	Name    string `yaml:"name"`
	Data    Data   `yaml:"data"`
}
*/
package yaml

import (
	"io/ioutil"

	"github.com/alexshnup/yaml.v2"
)

// ConfigManager struct
type ConfigManager struct {
	file string
}

// NewConfig constructor
func NewConfig(file string) *ConfigManager {
	return &ConfigManager{file}
}

// Load config from file
func (cm *ConfigManager) Load(config interface{}) error {
	data, err := ioutil.ReadFile(cm.file)
	if err != nil {
		return err
	}

	// config is a link
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return err
	}
	return nil
}
