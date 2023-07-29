package main

import (
	conf "clown-id/internal/config"
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

func main() {
	flag.Parse()
	config := parseConfig()
	print(config.BindAddr)
	// s := server.New(parseConfig())
	// err := s.Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
