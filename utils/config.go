package utils

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title    string
	Gameplay Gameplay
	Debug    Debug
	Server   Server
}

type Server struct {
	Host  string
	Port  string
	Admin string
}
type Gameplay struct {
	SizeField     [2]int
	CountResorces int
}

type Debug struct {
	LogLevel    string
	LogFilePath string
}

func NewConfig() *Config {
	var conf Config

	// ========= PARSE CONFIG ========= //
	confFile, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatalf("%v", err)
	}
	if err != nil {
		log.Fatalf("%v", err)
	}
	if _, err := toml.Decode(string(confFile), &conf); err != nil {
		log.Fatalf("%v", err)
	}

	return &conf

}
