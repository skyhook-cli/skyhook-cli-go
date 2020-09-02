package model

import (
	"fmt"
	"io/ioutil"

	"github.com/logrusorgru/aurora"

	"gopkg.in/yaml.v2"
)

const filename string = "./Tetherfile"

/*
Config represents the user's input
*/
type Config struct {
	ProjectType string            `yaml:"projectType"`
	Parameters  map[string]string `yaml:"parameters"`
}

/*
SaveConfig writes the config to a yaml file
*/
func (c Config) SaveConfig() error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0664)
}

/*
ReadConfig reads the yaml file into the config struct
*/
func (c *Config) ReadConfig() error {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		c.ProjectType = ""
		c.Parameters = make(map[string]string)
		fmt.Println(aurora.Yellow("Warning: Tetherfile not found, creating new one"))
		return err
	}

	err = yaml.Unmarshal(bytes, c)

	if err != nil {
		c.ProjectType = ""
		c.Parameters = make(map[string]string)
		fmt.Println(aurora.Yellow("Warning: there was an error reading the Tetherfile, creating new one"))
		return err
	}

	return nil
}
