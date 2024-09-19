package auth

import (
	"database/sql"
	"fmt"

	"todo-app/internal/common"
)

type Store struct {
	common.Store
	Table string
	db    *sql.DB
}

func NewStore() (*Store, error) {
	conn, err := common.OpenConnection()
	return &Store{db: conn, Table: "users"}, err
}

func (s *Store) RegisterUser(new RegisterUserDTO) (resp LoginResponseDTO, err error) {

	sql := fmt.Sprintf("INSERT INTO %s (name, email, password) VALUES (?, ?, ?) RETURNING *", s.Table)

	err = s.db.QueryRow(sql, new.Name, new.Email, new.Password).Scan(&resp.Name, &resp.ID, &resp.Email)

	return
}
