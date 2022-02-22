package utils

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type DefaultClient struct {
	Host string `toml:"Host"`
	Port string `toml:"Port"`
}

type ConfigClient struct {
	Server DefaultClient `toml:"Server"`
}

func GetClientConfig(path string) *ConfigClient {
	confFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var conf ConfigClient

	if _, err := toml.Decode(string(confFile), &conf); err != nil {
		log.Fatalf("%v", err)
	}

	return &conf
}
