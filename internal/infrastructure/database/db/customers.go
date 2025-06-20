package db

import (
	"context"
	"fmt"
	"miniapp/internal/dto"
)

func (r *repository) CreateNewCustomer(ctx context.Context, newCustomer dto.DBCreateCustomerDTO) (id int, err error) {
	// r.logger.Infoln("saving new customer in db")
	query := `
			INSERT INTO customers (
				tg_name,
				tg_id,
				phone_number
			)
			VALUES ($1, $2, 'none')
			RETURNING id
			`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, newCustomer.Name, newCustomer.TgId).Scan(&id)
	if id == 0 {
		return id, fmt.Errorf("failed to create customer")
	}
	// r.logger.Infoln("customer saved with id:", id)
	return id, nil
}

func (r *repository) LoadCustomer(ctx context.Context, id int) (customer dto.DBLoadCustomerDTO, err error) {
	// r.logger.Infoln("loading customer with tg_id:", tg_id)
	query := ` 
		SELECT 
			id, 
			tg_name, 
			tg_id, 
			phone_number, 
			created_at
		FROM customers 
		WHERE id = $1
	`
	// r.logger.Traceln("SQL Query:", formatQuery(query))
	row := r.client.QueryRow(ctx, query, id)
	err = row.Scan(&customer.Id, &customer.Name, &customer.TgId, &customer.PhoneNumber, &customer.CreatedAt)
	if err != nil {
		return customer, err
	}
	// r.logger.Infoln("customer loaded")
	return customer, nil
}

func (r *repository) LoadCustomerList(ctx context.Context) ([]dto.DBLoadCustomerDTO, error) {
	query := `
		SELECT 
			id, 
			tg_name, 
			tg_id, 
			phone_number, 
			created_at
		FROM customers
	`

	rows, err := r.client.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	customers := make([]dto.DBLoadCustomerDTO, 0)

	for rows.Next() {
		tempCustomer := dto.DBLoadCustomerDTO{}

		err = rows.Scan(&tempCustomer.Id, &tempCustomer.Name, &tempCustomer.TgId, &tempCustomer.PhoneNumber, &tempCustomer.CreatedAt)
		if err != nil {
			return nil, err
		}

		customers = append(customers, tempCustomer)
	}

	return customers, nil
}

func (r *repository) UpdateCustomer(ctx context.Context, customer dto.DBUpdateCustomerDTO) (id int, err error) {
	query := `
			UPDATE customers
			SET
				tg_name = $2,
				phone_number = $3
			WHERE id = $1
			RETURNING id
			`
	// r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, customer.Id, customer.Name, customer.PhoneNumber).Scan(&id)
	if id == 0 {
		return id, fmt.Errorf("failed to update customer")
	}
	// r.logger.Infoln("successful updated customer, id:", id)
	return id, nil
}
