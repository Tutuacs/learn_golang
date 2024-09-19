package auth

import (
	"database/sql"

	"todo-app/internal/common"
)

type Store struct {
	common.Store
	db *sql.DB
}

func NewStore() (*Store, error) {
	conn, err := common.OpenConnection()
	return &Store{db: conn}, err
}

func (s *Store) GetConn() *sql.DB {
	return s.db
}

func NewStoreInstace(conn *sql.DB) (*Store, error) {
	return &Store{db: conn}, nil
}

func (s *Store) CloseStore() {
	common.CloseConnection(s.db)
}

func (s *Store) RegisterUser(new RegisterUserDTO) (resp LoginResponseDTO, err error) {

	sql := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"

	result, err := s.db.Exec(sql, new.Name, new.Email, new.Password)
	if err != nil {
		return
	}

	resp.User.ID, err = result.LastInsertId()

	return
}
