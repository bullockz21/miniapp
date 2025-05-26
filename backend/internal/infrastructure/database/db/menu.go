package db

import (
	"context"
	"fmt"
	"miniapp/internal/dto"
)

// id SERIAL PRIMARY KEY,
// type_id INT NOT NULL,
// name VARCHAR(50) UNIQUE NOT NULL,
// price NUMERIC,
// description VARCHAR(1000)

func (r *repository) SaveNewMenuPosition(ctx context.Context, menu dto.MenuDTO) (id int, err error) {
	r.logger.Infoln("saving new menu position")
	query := `
			INSERT INTO menu (
				type_id,
				name,
				price,
				description
			)
			VALUES ($1, $2, $3, $4)
			RETURNING id
			`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, menu.Category, menu.Name, menu.Price, menu.Description).Scan(&id)
	if id == 0 {
		return id, fmt.Errorf("failed to create menu position")
	}
	r.logger.Infoln("created new menu position with id:", id)
	return id, nil
}

func (r *repository) LoadMenuPosition(ctx context.Context, id int) (menu dto.MenuDTO, err error) {
	r.logger.Infoln("loading menu position with id:", id)
	query := ` 
		SELECT 
			id,
			type_id,
			name,
			price,
			description
		FROM menu 
		WHERE id = $1
		`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	row := r.client.QueryRow(ctx, query, id)
	err = row.Scan(&menu.Id, &menu.Category, &menu.Name, &menu.Price, &menu.Description)
	if err != nil {
		return menu, err
	}
	return menu, nil
}

func (r *repository) LoadAllMenu(ctx context.Context) (menu []dto.MenuDTO, err error) {
	r.logger.Infoln("loading menu list")
	query := ` 
		SELECT 
			id,
			type_id,
			name,
			price,
			description
		FROM menu 
		`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	menuList := make([]dto.MenuDTO, 0)
	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tempMenu := dto.MenuDTO{}
		err = rows.Scan(&tempMenu.Id, &tempMenu.Category, &tempMenu.Name, &tempMenu.Price, &tempMenu.Description)
		if err != nil {
			return nil, err
		}
		menuList = append(menuList, tempMenu)
	}
	return menuList, nil
}

func (r *repository) UpdateMenuPosition(ctx context.Context, menu dto.MenuDTO) (id int, err error) {
	r.logger.Infoln("updating menu position with id:", menu.Id)
	query := `
			UPDATE menu
			SET
				type_id,
				name,
				price,
				description
			WHERE id = $1
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
			`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, menu.Id, menu.Category, menu.Name, menu.Price, menu.Description).Scan(&id)
	if id == 0 {
		return id, fmt.Errorf("failed to update menu position")
	}
	r.logger.Infoln("successful updated menu position, id:", id)
	return id, nil
}

func (r *repository) DeleteMenuPosition(ctx context.Context, id int) error {
	ID := 0
	query := `
		DELETE FROM menu 
		WHERE id = $1
		RETURNING id
		`
	r.client.QueryRow(ctx, query, id).Scan(&ID)
	r.logger.Traceln("SQL Query:", formatQuery(query))
	if ID == 0 {
		return fmt.Errorf("menu position not found")
	}
	return nil
}
