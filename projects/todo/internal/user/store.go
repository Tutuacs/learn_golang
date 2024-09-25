package user

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

func (s *Store) GetConn() *sql.DB {
	return s.db
}

func NewStoreInstance(conn *sql.DB) (*Store, error) {
	return &Store{db: conn, Table: "users"}, nil
}

func (s *Store) CloseStore() {
	common.CloseConnection(s.db)
}

func (s *Store) CreateUser() {

}

func (s *Store) ListUsers() (users []User, err error) {

	sql := fmt.Sprintf("SELECT * FROM %s", s.Table)

	rows, err := s.db.Query(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		user := rowsToUser(rows)
		if user.ID != 0 {
			users = append(users, user)
		}
	}

	return
}

func (s *Store) FindUserByID(id int64) (user User, err error) {

	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", s.Table)

	row := s.db.QueryRow(sql, id)

	user = rowToUser(row)

	return
}

func (s *Store) FindUserByEmail(email string) (user User, err error) {

	sql := fmt.Sprintf("SELECT * FROM %s WHERE email = ?", s.Table)

	row := s.db.QueryRow(sql, email)

	user = rowToUser(row)

	return
}

func rowToUser(row *sql.Row) (user User) {

	_ = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	return
}

func rowsToUser(row *sql.Rows) (user User) {

	_ = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	return
}
