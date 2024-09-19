package user

import "time"

var Enum = map[string]int{
	"ROLE_CLIENT": 0,
	"ROLE_ADMIN":  1,
}

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      int       `json:"role" validate:"required,min=0,max=1"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Role struct {
}
