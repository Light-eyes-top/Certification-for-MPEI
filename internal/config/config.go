package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

type Postgres struct {
	User     string `yaml:"user"`
	Port     string `yaml:"port"`
	Password string `yaml:"pass"`
	Host     string `yaml:"host"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
}

type Server struct {
	PortREST string `yaml:"portREST"`
}

type Config struct {
	Postgres Postgres `yaml:"psql"`
	Server   Server   `yaml:"server"`
}

func Init() (*Config, error) {
	filePath := flag.String("c", "config.yml", "Path to configuration file")
	flag.Parse()
	config := &Config{}
	data, err := os.ReadFile(*filePath)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return config, nil
}
