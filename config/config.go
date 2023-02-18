package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port     string `yaml:"port"`
	Database Database
}

type Database struct {
	Host     string `yaml:"host"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbport   int    `yaml:"dbport"`
	Sslmode  string `yaml:"sslmode"`
	Tz       string `yaml:"tz"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ReadConfigYaml(filename string, destination *Config) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("file not found %v", err)
		return
	}

	err = yaml.Unmarshal(b, destination)
	if err != nil {
		log.Fatalf("error not convert to struct %v", err)
		return
	}
}
