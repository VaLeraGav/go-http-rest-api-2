package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"local"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"ADDRESS" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"60s"`
	User        string        `yaml:"user" env:"USER" env-required:"true"`
	Password    string        `yaml:"password" env:"HTTP_SERVER_PASSWORD" env-required:"true"`
}

func MustLoad() *Config {
	// configPath := os.Getenv("CONFIG_PATH")
	configPath := "./config/prod.yaml"
	absoluteConfigPath, err := filepath.Abs(configPath)
	if err != nil {
		log.Fatalf("Error in getting an absolute path: %v", err)
	}

	if absoluteConfigPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(absoluteConfigPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", absoluteConfigPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(absoluteConfigPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
