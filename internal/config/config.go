package config

import "github.com/sirupsen/logrus"

type Config struct {
	Port         int          `toml:"port"`
	LogLevel     logrus.Level `toml:"log_level"`
	EmaClientURL string       `toml:"ema-client-url"`
}

func NewConfig() *Config {
	return &Config{
		Port:         8080,
		LogLevel:     logrus.DebugLevel,
		EmaClientURL: "http://localhost:8081/",
	}
}
