package todo

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
	return &Store{db: conn, Table: "todos"}, err
}

func (s *Store) CloseStore() {
	common.CloseConnection(s.db)
}

func (s *Store) CreateTodo(new NewTodoDTO) (todo Todo, err error) {

	sql := fmt.Sprintf("INSERT INTO %s (name, done, userID) VALUES (?, ?, ?) RETURNING *", s.Table)

	err = s.db.QueryRow(sql, new.Name, new.Done, new.UserID).Scan(&todo.ID, &todo.Name, &todo.Done, &todo.UserID)

	return
}
