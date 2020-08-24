package model

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const filename string = "./Tetherfile"

/*
Config represents the user's input
*/
type Config interface {
	SaveConfig() error
	ReadConfig() (Config, error)
}

/*
AppParams are config options specific to an app project
*/
type AppParams struct {
	ContainerPlatform string `yaml:"containerPlatform"`
	DockerImage       string `yaml:"dockerImage"`
	AppName           string `yaml:"appName"`
}

/*
AppConfig is a concrete Config for app projects
*/
type AppConfig struct {
	InfraOrApp string    `yaml:"infraOrApp"`
	AppParams  AppParams `yaml:"appParams"`
}

/*
InfraParams are config options specific to an infra project
*/
type InfraParams struct {
	ProjectName           string `yaml:"projectName"`
	ContainerPlatform     string `yaml:"containerPlatform"`
	ImageRepository       string `yaml:"imageRepository"`
	JenkinsAdminUser      string `yaml:"jenkinsAdminUser"`
	JenkinsAdminPassword  string `yaml:"jenkinsAdminPassword"`
	SonarAdminUser        string `yaml:"sonarAdminUser"`
	SonarAdminPassword    string `yaml:"sonarAdminPassword"`
	AnchoreAdminUsername  string `yaml:"anchoreAdminUsername"`
	AnchoreAdminPassword  string `yaml:"anchoreAdminPassword"`
	SeleniumAdminUsername string `yaml:"seleniumAdminUsername"`
	SeleniumAdminPassword string `yaml:"seleniumAdminPassword"`
}

/*
InfraConfig is a concrete Config for infra projects
*/
type InfraConfig struct {
	InfraOrApp  string      `yaml:"infraOrApp"`
	InfraParams InfraParams `yaml:"infraParams"`
}

/*
SaveConfig writes the config to a yaml file
*/
func (c AppConfig) SaveConfig() error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0664)
}

/*
ReadConfig reads the yaml file into the config struct
*/
func (c AppConfig) ReadConfig() (Config, error) {
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

/*
SaveConfig writes the config to a yaml file
*/
func (c InfraConfig) SaveConfig() error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0664)
}

/*
ReadConfig reads the yaml file into the config struct
*/
func (c InfraConfig) ReadConfig() (Config, error) {
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
