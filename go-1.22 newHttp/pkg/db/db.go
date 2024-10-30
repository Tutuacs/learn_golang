package database

import (
	"database/sql"
	"fmt"

	"go-1.22-http/pkg/config"
)

type Store struct {
}

func (s *Store) OpenConnection() (conn *sql.DB, err error) {
	conf := config.GetDbConfig()

	stringConnection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Name)

	conn, err = sql.Open("postgres", stringConnection)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()

	return
}

func (s *Store) CloseConnection(conn *sql.DB) error {
	return conn.Close()
}
