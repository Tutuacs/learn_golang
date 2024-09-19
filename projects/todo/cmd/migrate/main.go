package main

import (
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"todo-app/internal/common"
)

func main() {

	conn, err := common.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(conn, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	} else if cmd == "down" {
		if err := migration.Down(); err != nil && err != migrate.ErrNoChange {
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
		if err := migration.Force(version); err != nil {
			log.Fatal(err)
		}
	}

}
