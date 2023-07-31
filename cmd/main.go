package main

import (
	conf "clown-id/internal/config"
	"clown-id/internal/server"
	"encoding/json"
	"flag"
	"log"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.json", "path to config file")
}

func parseConfig() *conf.Config {
	config := conf.NewConfig()
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("Could not open config file: ", err)
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		log.Fatal("Could not parse config file: ", err)
	}
	return config
}

// @title			Clown-ID API
// @version			0.1
// @description		Auth service for clown-devs projects
func main() {
	flag.Parse()

	s := server.New(parseConfig())
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
