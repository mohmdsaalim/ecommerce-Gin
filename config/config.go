package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Port string `yaml:"port"`
	} `yaml:"app"`

	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`

	Postgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"postgres"`

	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`

	SMTP struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Email    string `yaml:"email"`
		Password string `yaml:"password"`
	} `yaml:"smtp"`
}

var AppConfig *Config // globally accessable

func LoadConfig() {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalf("[error] failed to open config.yaml: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	AppConfig = &Config{}
	if err := decoder.Decode(AppConfig); err != nil {
		log.Fatalf("[error] failed to decode YAML: %v", err)
	}

	log.Println("✅ Config loaded from config.yaml")
}