package configs

import (
	"errors"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Port     string
	Database string
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	if host == "" || username == "" || password == "" || port == "" || database == "" {
		return nil, errors.New("database config not found")
	}

	if _port := port; _port == "" {
		port = "5000"
	}

	return &DatabaseConfig{
		Host:     host,
		Username: username,
		Password: password,
		Port:     port,
		Database: database,
	}, nil
}
