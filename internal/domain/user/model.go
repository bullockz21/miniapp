package user

import "time"

// id SERIAL PRIMARY KEY,
// user_name VARCHAR(150) UNIQUE NOT NULL,
// role_id INT DEFAULT 0,
// password_hash VARCHAR(100),
// email VARCHAR(100),
// created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// deleted_at TIMESTAMP

type User struct {
	Id           int       `json:"id"`
	UserName     string    `json:"user_name"`
	RoleId       int       `json:"role_id"`
	Password     string    `json:"password"`
	PasswordHash string    `json:"password_hash"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
