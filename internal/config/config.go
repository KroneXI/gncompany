package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config конвертирован из config.yaml в структуру
type Config struct {
	Server `yaml:"Server"`
}

type User struct {
	Groups      []Group `yaml:"Groups"`
	PhoneNumber string  `yaml:"PhoneNumber"`
}

type Group struct {
	Messenger string `yaml:"Messenger"`
	URL       string `yaml:"URL"`
	ImageURL  string `yaml:"ImageURL"`
}

type AdminAUth struct {
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

type Server struct {
	Port      string `yaml:"Port"`
	User      `yaml:"User"`
	AdminAUth `yaml:"AdminAUth"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
