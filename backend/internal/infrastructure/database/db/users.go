package db

import (
	"context"
	"fmt"
	"miniapp/internal/dto"
)

// id SERIAL PRIMARY KEY,
// user_name VARCHAR(150) UNIQUE, NOT NULL
// role_id INT,
// password_hash VARCHAR(100),
// email VARCHAR(100),
// created_at DATE,
// deleted_at DATE

func (r *repository) SaveNewUser(ctx context.Context, newUser dto.UserDTO) (id int, err error) {

	r.logger.Infoln("creating user")

	query := `
			INSERT INTO users (
				user_name,
				password_hash,
				email
			)
			VALUES ($1, $2, $3)
			RETURNING id
			`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, newUser.UserName, newUser.PasswordHash, newUser.Email).Scan(&id)
	if id == 0 {
		return id, fmt.Errorf("failed to create user")
	}

	r.logger.Infoln("created user with id:", id)

	return id, nil
}

func (r *repository) LoadUser(ctx context.Context, id int) (user dto.UserDTO, err error) {

	r.logger.Infoln("loading user with id:", id)

	query := ` 
		SELECT 
			id,
			user_name,
			role_id,
			password_hash,
			email,
			created_at,
			deleted_at
		FROM customers 
		WHERE id = $1
		`

	r.logger.Traceln("SQL Query:", formatQuery(query))
	row := r.client.QueryRow(ctx, query, id)
	err = row.Scan(&user.Id, &user.UserName, &user.RoleId, &user.PasswordHash, &user.Email, &user.CreatedAt, &user.DeletedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateUser(ctx context.Context, user dto.UserDTO) (id int, err error) {

	r.logger.Infoln("updating customer, id:", user.Id)

	query := `
			UPDATE users
			SET
				user_name,
				role_id,
				password_hash,
				email
			WHERE id = $1
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
			`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, user.UserName, user.RoleId, user.PasswordHash, user.Email).Scan(&id)
	if id == 0 {
		return id, fmt.Errorf("failed to update user")
	}

	r.logger.Infoln("successful updated user, id:", id)

	return id, nil
}

func (r *repository) DeleteUser(ctx context.Context, id int) (err error) {

	r.logger.Infoln("updating customer, id:", id)

	query := `
			UPDATE users
			SET
				role_id,
				deleted_at
			WHERE id = $1
			VALUES ($1, 5, TIMESTAMP)
			RETURNING id
			`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, id).Scan(&id)
	if id == 0 {
		return fmt.Errorf("failed to delete user")
	}

	r.logger.Infoln("successful updated user, id:", id)

	return nil
}
