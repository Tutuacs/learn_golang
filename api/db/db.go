package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(string(Red), err, Reset)
	}

	return db, err
}

func InitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(string(Red), err, Reset)
	}

	log.Println(string(Green), "DB: Successfully connected!", Reset)
}
