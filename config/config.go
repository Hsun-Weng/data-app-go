package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Profile string `yaml:"profile"`
	Server  struct {
		Port string `yaml:"port"`
	}
	Mongodb struct {
		Host                   string `yaml:"host"`
		Database               string `yaml:"database"`
		UserName               string `yaml:"username"`
		Password               string `yaml:"password"`
		AuthenticationDatabase string `yaml:"authentication-database"`
	}
}

func ReadConfig(configPath string) *Config {
	config := new(Config)
	configFile, err := ioutil.ReadFile(configPath)

	log.Println("configFile:", configFile)

	if err != nil {
		log.Printf("configFile.Get err #%v", err)
	}
	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("config", config)
	return config
}
