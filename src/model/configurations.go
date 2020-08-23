package model

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const filename string = "./Tetherfile"

type Config interface {
	SaveConfig() error
	ReadConfig() (Config, error)
}

type AppParams struct {
	ContainerPlatform string `yaml:"containerPlatform"`
	DockerImage       string `yaml:"dockerImage"`
	AppName           string `yaml:"appName"`
}

type AppConfig struct {
	InfraOrApp string    `yaml:"infraOrApp"`
	AppParams  AppParams `yaml:"appParams"`
}

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

type InfraConfig struct {
	InfraOrApp  string      `yaml:"infraOrApp"`
	InfraParams InfraParams `yaml:"infraParams"`
}

func (c AppConfig) SaveConfig() error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0664)
}

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

func (c InfraConfig) SaveConfig() error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0664)
}

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
