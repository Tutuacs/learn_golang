package common

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/database/mysql"

	"todo-app/pkg/config"
	"todo-app/pkg/logs"
)

type Store struct {
}

func OpenConnection() (*sql.DB, error) {

	conf := config.GetDBConfig()

	dbConfig := mysql.Config{
		User:                 conf.DB_User,
		Passwd:               conf.DB_Pass,
		Addr:                 conf.DB_Addr,
		DBName:               conf.DB_Name,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		logs.ErrorLog(fmt.Sprintf("Error opening B connection: %v", err))
	}

	return db, err

}
