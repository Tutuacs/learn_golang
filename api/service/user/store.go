package user

import (
	"database/sql"
	"fmt"

	"github.com/Tutuacs/learn_golang/api.git/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {

	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	usr := new(types.User)
	for rows.Next() {
		usr, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if usr.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return usr, nil
}

func (s *Store) GetUserByID(ID int) (*types.User, error) {

	sql := "SELECT * FROM users WHERE id = ?"

	rows, err := s.db.Query(sql, ID)
	if err != nil {
		return nil, err
	}

	usr := new(types.User)

	for rows.Next() {
		usr, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	return usr, err
}

func (s *Store) CreateUser(user types.User) error {

	sql := "INSERT INTO users (firstName, lastName, email, password) VALUES (?,?,?,?)"
	_, err := s.db.Exec(sql, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
