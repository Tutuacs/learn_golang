package main

import (
	"log"
	"os"
	"strconv"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/Tutuacs/learn_golang/api.git/config"
	"github.com/Tutuacs/learn_golang/api.git/db"
)

func main() {
	dbConfig := mysqlCfg.Config{
		User:                 config.Envs.DbUser,
		Passwd:               config.Envs.DbPass,
		Addr:                 config.Envs.DbAdrr,
		DBName:               config.Envs.DbName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	dbConnection, err := db.NewMySQLStorage(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(dbConnection, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	} else if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	} else if cmd == "force" {
		if len(os.Args) < 3 {
			log.Fatal("Please provide a version number to force")
		}
		version, err := strconv.Atoi(os.Args[len(os.Args)-2])
		if err != nil {
			log.Fatal("Invalid version number")
		}
		if err := m.Force(version); err != nil {
			log.Fatal(err)
		}
	}

}
