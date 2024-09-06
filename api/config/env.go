package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiConfig
	DbConfig
	JWT_EXP    int64
	JWT_SECRET string
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
		JWT_EXP:    getEnvAsInt("JWT_EXP", 3600*24*7),
		JWT_SECRET: getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
