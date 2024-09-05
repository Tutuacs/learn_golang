package main

import (
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/Tutuacs/learn_golang/api.git/cmd/api"
	"github.com/Tutuacs/learn_golang/api.git/config"
	"github.com/Tutuacs/learn_golang/api.git/db"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func main() {

	dbConnection, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DbUser,
		Passwd:               config.Envs.DbPass,
		Addr:                 config.Envs.DbAdrr,
		DBName:               config.Envs.DbName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(string(Red), err, Reset)
	}

	db.InitStorage(dbConnection)

	server := api.NewApiServer(config.Envs.ApiPort, dbConnection)
	if err := server.Run(); err != nil {
		log.Fatal(string(Red), err, Reset)
	}

}
