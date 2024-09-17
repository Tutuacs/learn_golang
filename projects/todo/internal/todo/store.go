package todo

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

func (s *Store) CloseStore() {
	common.CloseConnection(s.db)
}

func (s *Store) CreateTodo(todo Todo) error {

	return nil
}
