package db

import (
	"context"
	"errors"
	"fmt"
	"miniapp/internal/dto"

	"github.com/jackc/pgx/v5/pgconn"
)

// id SERIAL PRIMARY KEY,
// user_name VARCHAR(150) UNIQUE, NOT NULL
// role_id INT,
// password_hash VARCHAR(100),
// email VARCHAR(100),
// created_at DATE,
// deleted_at DATE

func (r *repository) CreateNewUser(ctx context.Context, DBuser dto.DBCreateUserDTO) (id int, err error) {
	// r.logger.Infoln("creating user")
	query := `
			INSERT INTO users (
				user_name,
				password_hash,
				email
			)
			VALUES ($1, $2, $3)
			RETURNING id
			`
	// r.logger.Traceln("SQL Query:", formatQuery(query))
	err = r.client.QueryRow(ctx, query, DBuser.Name, DBuser.PasswordHash, DBuser.Email).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return id, fmt.Errorf("code: %s, err: %s", pgErr.Code, pgErr.Message)
		}
	}
	// r.logger.Infoln("created user with id:", id)
	return id, nil
}

// func (r *repository) LoadUser(ctx context.Context, id int) (user dto.UserDTO, err error) {
// 	r.logger.Infoln("loading user with id:", id)
// 	query := `
// 		SELECT
// 			id,
// 			user_name,
// 			role_id,
// 			password_hash,
// 			email,
// 			created_at,
// 			deleted_at
// 		FROM users
// 		WHERE id = $1
// 		`
// 	r.logger.Traceln("SQL Query:", formatQuery(query))
// 	row := r.client.QueryRow(ctx, query, id)
// 	err = row.Scan(&user.Id, &user.UserName, &user.RoleId, &user.PasswordHash, &user.Email, &user.CreatedAt, &user.DeletedAt)
// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

// func (r *repository) LoadAllUsers(ctx context.Context) (users []dto.UserDTO, err error) {
// 	r.logger.Infoln("loading all users")
// 	query := `
// 		SELECT
// 			id,
// 			user_name,
// 			role_id,
// 			password_hash,
// 			email,
// 			created_at,
// 			deleted_at
// 		FROM users
// 		`
// 	r.logger.Traceln("SQL Query:", formatQuery(query))
// 	rows, err := r.client.Query(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		tempUser := dto.UserDTO{}
// 		err = rows.Scan(&tempUser.Id, &tempUser.UserName, &tempUser.RoleId, &tempUser.PasswordHash, &tempUser.Email, &tempUser.CreatedAt, &tempUser.DeletedAt)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, tempUser)
// 	}
// 	return users, nil
// }

// func (r *repository) UpdateUserPassword(ctx context.Context, user dto.UserDTO) (id int, err error) {
// 	r.logger.Infoln("updating users password, id:", user.Id)
// 	query := `
// 			UPDATE users
// 			SET
// 				password_hash = $2
// 			WHERE id = $1
// 			RETURNING id
// 			`
// 	r.logger.Traceln("SQL Query:", formatQuery(query))
// 	r.client.QueryRow(ctx, query, user.Id, user.PasswordHash).Scan(&id)
// 	if id == 0 {
// 		return id, fmt.Errorf("failed to update user")
// 	}
// 	r.logger.Infoln("successful updated user, id:", id)
// 	return id, nil
// }

// func (r *repository) UpdateUserEmail(ctx context.Context, user dto.UserDTO) (id int, err error) {
// 	r.logger.Infoln("updating users email, id:", user.Id)
// 	query := `
// 			UPDATE users
// 			SET
// 				email = $1
// 			WHERE id = $2
// 			RETURNING id
// 			`
// 	r.logger.Traceln("SQL Query:", formatQuery(query))
// 	r.client.QueryRow(ctx, query, user.Email, user.Id).Scan(&id)
// 	if id == 0 {
// 		return id, fmt.Errorf("failed to update user")
// 	}
// 	r.logger.Infoln("successful updated user, id:", id)
// 	return id, nil
// }

// func (r *repository) UpdateUserRoleId(ctx context.Context, user dto.UserDTO) (id int, err error) {
// 	r.logger.Infoln("updating users role id, id:", user.Id)
// 	query := `
// 			UPDATE users
// 			SET
// 				role_id = $2
// 			WHERE id = $1
// 			RETURNING id
// 			`
// 	r.logger.Traceln("SQL Query:", formatQuery(query))
// 	r.client.QueryRow(ctx, query, user.Id, user.RoleId).Scan(&id)
// 	if id == 0 {
// 		return id, fmt.Errorf("failed to update user")
// 	}
// 	r.logger.Infoln("successful updated user, id:", id)
// 	return id, nil
// }

// func (r *repository) DeleteUser(ctx context.Context, id int) (err error) {
// 	r.logger.Infoln("updating customer, id:", id)
// 	query := `
// 			UPDATE users
// 			SET
// 				role_id,
// 				deleted_at
// 			WHERE id = $1
// 			VALUES ($1, 5, TIMESTAMP)
// 			RETURNING id
// 			`
// 	r.logger.Traceln("SQL Query:", formatQuery(query))
// 	r.client.QueryRow(ctx, query, id).Scan(&id)
// 	if id == 0 {
// 		return fmt.Errorf("failed to delete user")
// 	}
// 	r.logger.Infoln("successful updated user, id:", id)
// 	return nil
// }
