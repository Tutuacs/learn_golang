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

func NewStoreInstance(db *sql.DB) (*Store, error) {
	return &Store{db: db, Table: "todos"}, nil
}

func (s *Store) CreateTodo(new NewTodoDTO) (todo Todo, err error) {

	sql := fmt.Sprintf("INSERT INTO %s (name, done, userID) VALUES (?, ?, ?) RETURNING *", s.Table)

	row := s.db.QueryRow(sql, new.Name, new.Done, new.UserID)

	todo = rowToTodo(row)

	return
}

func (s *Store) FindTodoByID(id int64) (todo Todo, err error) {

	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", s.Table)

	row := s.db.QueryRow(sql, id)

	todo = rowToTodo(row)

	return
}

func (s *Store) ListTodo() (todos []Todo, err error) {

	sql := fmt.Sprintf("SELECT * FROM %s", s.Table)

	rows, err := s.db.Query(sql)

	for rows.Next() {
		todo := rowsToTodo(rows)
		if todo.ID != 0 {
			todos = append(todos, todo)
		}
	}

	return
}

func (s *Store) ListTodosByUserID(id int64) (todos []Todo, err error) {

	sql := fmt.Sprintf("SELECT * FROM %s WHERE userId = ?", s.Table)

	rows, err := s.db.Query(sql, id)

	for rows.Next() {
		todo := rowsToTodo(rows)
		if todo.ID != 0 {
			todos = append(todos, todo)
		}
	}

	return
}

func rowsToTodo(rows *sql.Rows) (todo Todo) {

	_ = rows.Scan(&todo.ID, &todo.Name, &todo.Done, &todo.UserID)

	return
}

func rowToTodo(row *sql.Row) (todo Todo) {

	_ = row.Scan(&todo.ID, &todo.Name, &todo.Done, &todo.UserID)

	return
}
