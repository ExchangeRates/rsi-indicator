package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"rsi_indicator/internal"
	"rsi_indicator/internal/config"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/service.toml", "path to config file")
}

func main() {
	flag.Parse()

	configuration := config.NewConfig()
	if _, err := toml.DecodeFile(configPath, configuration); err != nil {
		log.Fatal(err)
	}

	if err := internal.Start(configuration); err != nil {
		log.Fatal(err)
	}

}
