package main

import (
	"fmt"

	"todo-app/cmd/api"
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

	apiServer := api.NewApiServer(apiConf.API_Port)

	if err := apiServer.Start(); err != nil {
		logs.ErrorLog("FailsRaning Server")
		return
	}

}
