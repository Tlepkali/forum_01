package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	HashedPW  string    `json:"hashed_pw"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UpdateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginUserDTO struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

type UserService interface {
	CreateUser(user *CreateUserDTO) error
	GetUserByID(id int) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *UpdateUserDTO) error
	LoginUser(user *LoginUserDTO) (int, error)
}

type UserRepo interface {
	CreateUser(user *User) error
	GetUserByID(id int) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
}
