package todo

import (
	"database/sql"

	"todo-app/internal/common"
)

type Store struct {
	common.Store
	db *sql.DB
}

func NewStore() (Store, error) {
	newStore, err := common.OpenConnection()

	return Store{db: newStore}, err
}

func (s *Store) UserCreate() {

	return
}
