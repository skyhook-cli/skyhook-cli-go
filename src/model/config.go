package model

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const filename string = "./Tetherfile"

/*
Config represents the user's input
*/
type Config struct {
	InfraOrApp string            `yaml:"infraOrApp"`
	Parameters map[string]string `yaml:"parameters"`
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
func (c Config) ReadConfig() (Config, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(bytes, &c)

	if err != nil {
		return c, err
	}

	return c, nil
}
