package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"todo-app/pkg/logs"
)

type apiConfig struct {
	API_Port string
}

type jwtConfig struct {
	JWT_Secret     string
	JWT_Expiration int64
}

type dbConfig struct {
	DB_Port string
	DB_Host string
	DB_Addr string
	DB_User string
	DB_Pass string
	DB_Name string
}

type config struct {
	APIConfig apiConfig
	JWTConfig jwtConfig
	DBConfig  dbConfig
}

var cfg *config

func init() {
	logs.MessageLog("Initializing config variables...")
	godotenv.Load()
	cfg = new(config)
	cfg = defaultConfig()
}

func defaultConfig() *config {
	return &config{
		APIConfig: apiConfig{
			API_Port: getEnv("API_PORT", ":9000"),
		},
		JWTConfig: jwtConfig{
			JWT_Secret:     getEnv("JWT_SECRET", "Secret"),
			JWT_Expiration: getNumberEnv("JWT_EXPIRATION", 3600*24*7),
		},
		DBConfig: dbConfig{
			DB_Port: getEnv("DB_PORT", "9999"),
			DB_Host: getEnv("DB_HOST", "127.0.0.1"),
			DB_Addr: getEnv("DB_ADDR", "127.0.0.1:9999"),
			DB_User: getEnv("DB_USER", "todo_user"),
			DB_Pass: getEnv("DB_PASS", "todo_pass"),
			DB_Name: getEnv("DB_NAME", "todo-app"),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok || len(value) == 0 {
		logs.WarnLog(fmt.Sprintf("the value of %s on .ENV file is invalid, value: '%s', using the default value to '%s':'%s'", key, value, key, defaultValue))
		return defaultValue
	}
	return value
}

func getNumberEnv(key string, defaultValue int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return intValue
		}
		logs.WarnLog(fmt.Sprintf("the value of %s on .ENV file is invalid, value: '%s', using the default value to '%s':'%d'", key, value, key, defaultValue))
	}
	return defaultValue
}

func GetDBConfig() dbConfig {
	return cfg.DBConfig
}

func GetApiConfig() apiConfig {
	return cfg.APIConfig
}

func GetJWTConfig() jwtConfig {
	return cfg.JWTConfig
}
