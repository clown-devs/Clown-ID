package config

import "github.com/sirupsen/logrus"

type Config struct {
	BindAddr       string `json:"bindAddr"`
	LogLevel       string `json:"logLevel"`
	DbConnStr      string `json:"dbConnStr"`
	Salt           string `json:"salt"`
	ApiPrefix      string `json:"apiPrefix"`
	SwaggerEnabled bool   `json:"swaggerEnabled"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:       ":8080",
		LogLevel:       logrus.DebugLevel.String(),
		Salt:           "salt",
		SwaggerEnabled: true,
	}
}
