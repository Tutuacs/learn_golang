package auth

import "todo-app/internal/user"

type RegisterUserDTO struct {
	LoginDTO
	Name string `json:"name"`
}

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=15"`
}

type LoginResponseDTO struct {
	Token string    `json:"token"`
	User  user.User `json:"user"`
}
