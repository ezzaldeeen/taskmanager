package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

const (
	cfgPath = "config.yml"
)

// Config represents the application configuration.
type Config struct {
	// Service configuration section.
	Service struct {
		Port string `yaml:"port" env:"PORT" env-default:"8080"`
	} `yaml:"service"`
	// DB configuration section.
	DB struct {
		Port     string `yaml:"port" env:"PORT" env-default:"5432"`
		Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
		Name     string `yaml:"name" env:"NAME" env-default:"postgres"`
		User     string `yaml:"user" env:"USER" env-default:"user"`
		Password string `yaml:"password" env:"PASSWORD"`
	} `yaml:"database"`
	Environment string
}

var (
	cfg  Config
	once sync.Once
)

// Load reads the configuration from a YAML file and returns
// the configuration instance. It ensures that the configuration is loaded
// only once using the sync.Once mechanism (singleton).
func Load() *Config {
	once.Do(func() {
		err := cleanenv.ReadConfig(cfgPath, &cfg)
		if err != nil {
			log.Fatal("Error loading configuration:", err)
		}
	})
	return &cfg
}
