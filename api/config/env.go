package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiConfig
	DbConfig
}

type DbConfig struct {
	DbUser string
	DbHost string
	DbPort string
	DbPass string
	DbAdrr string
	DbName string
}

type ApiConfig struct {
	ApiHost string
	ApiPort string
}

var Envs = initConfig()

func initConfig() *Config {
	godotenv.Load()
	return &Config{
		ApiConfig: ApiConfig{
			ApiHost: getEnv("API_HOST", "http://localhost"),
			ApiPort: getEnv("API_PORT", ":9000"),
		},
		DbConfig: DbConfig{
			DbHost: getEnv("DB_HOST", "127.0.1"),
			DbPort: getEnv("DB_PORT", ":3306"),
			DbAdrr: fmt.Sprintf("%s:%s",
				getEnv("DB_HOST", "127.0.1"),
				getEnv("DB_PORT", ":3306"),
			),
			DbUser: getEnv("DB_USER", "user_ecom"),
			DbPass: getEnv("DB_PASS", "ecom_password"),
			DbName: getEnv("DB_NAME", "ecom"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
