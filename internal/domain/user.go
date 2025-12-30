package domain

import "time"

type User struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	RoleID       int64     `json:"role_id"`
	Role         *Role     `json:"role,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Role struct {
	ID    int64  `json:"id"`
	Value string `json:"value"`
}

type UserRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
