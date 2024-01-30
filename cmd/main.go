package main

import (
	conf "clown-id/internal/config"
	"clown-id/internal/server"
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.json", "path to config file")
}

func parseConfig() *conf.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	config := &conf.Config{
		BindAddr:       os.Getenv("BIND_ADDR"),
		LogLevel:       os.Getenv("LOG_LEVEL"),
		DbConnStr:      os.Getenv("DB_CONN_STR"),
		MigrationStr:   os.Getenv("MIGRATION_STR"),
		Salt:           os.Getenv("SALT"),
		ApiPrefix:      os.Getenv("API_PREFIX"),
		SwaggerEnabled: os.Getenv("SWAGGER_ENABLED") == "true",
		Secret:         os.Getenv("SECRET"),
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
