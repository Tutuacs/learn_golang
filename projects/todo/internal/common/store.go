package common

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"

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

	conn, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		logs.ErrorLog(fmt.Sprintf("Error opening DB connection: %v", err))
	}

	return conn, err

}

func CloseConnection(conn *sql.DB) {
	conn.Close()
}
