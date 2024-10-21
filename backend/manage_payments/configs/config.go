package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type TConfig struct {
	Service  *ServiceConfig
	Database *DatabaseConfig
}

var Config TConfig

func LoadConfig() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file : " + err.Error())
	}

	database, err := LoadDatabaseConfig()

	if err != nil {
		log.Fatalln(err)
	}

	Config = TConfig{
		Service:  LoadServiceConfig(),
		Database: database,
	}
}

// func GetConfig(key string) string {
// 	return os.Getenv(key)
// }

func GetConfig() TConfig {
	return Config
}
