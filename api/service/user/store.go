package user

import (
	"database/sql"
	"fmt"
	"time"

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
	return &types.User{
		ID:        1,
		FirstName: "user_name",
		LastName:  "user_lastname",
		Email:     "email",
		Password:  "pass",
		CreatedAt: time.Time{},
	}, nil
}

func (s *Store) CreateUser(types.User) error {
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
