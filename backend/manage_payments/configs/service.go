package configs

import (
	"os"
)

type ServiceConfig struct {
	Env        string
	ServerPort string
	SSL        string
}

func LoadServiceConfig() *ServiceConfig {
	env := os.Getenv("ENV")
	port := os.Getenv("SERVER_PORT")
	ssl := os.Getenv("SSL")

	if env == "" {
		env = "development"
	}

	if _port := port; _port == "" {
		port = "5000"
	}

	return &ServiceConfig{
		Env:        env,
		ServerPort: port,
		SSL:        ssl,
	}
}
