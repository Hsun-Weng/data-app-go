package config

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Server struct{
		Port string `yaml:"port"`
	}
	Mongodb struct{
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Database string `yaml:"database"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
		AuthenticationDatabase string `yaml:"authentication-database"`
	}
}

func ReadConfig() *Config {
	config := new(Config)
	configFile, err := ioutil.ReadFile("config.yml")

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
