package dto

import "time"

type WebCreateUserDTO struct {
	Name     string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type DBCreateUserDTO struct {
	Id           int       `json:"id"`
	Name         string    `json:"user_name"`
	RoleId       int       `json:"role_id"`
	PasswordHash string    `json:"password_hash"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
