package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var cf *Config

func main() {
	fmt.Println(cf.Url)
}

// profile variables
type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
	Database string `yaml:"database"`
	Test     string `yaml:"test"`
}

func (c *Config) getConf() *Config {
	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
