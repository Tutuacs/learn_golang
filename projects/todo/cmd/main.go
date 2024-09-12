package main

import (
	"fmt"

	"todo-app/pkg/config"
	"todo-app/pkg/logs"
)

func main() {
	dbConf := config.GetDBConfig()
	apiConf := config.GetApiConfig()
	jwtConf := config.GetJWTConfig()

	logs.MessageLog(fmt.Sprintf("DB_CONF: %v", dbConf))
	logs.MessageLog(fmt.Sprintf("API_CONF: %v", apiConf))
	logs.MessageLog(fmt.Sprintf("JWT_CONF: %v", jwtConf))
}
