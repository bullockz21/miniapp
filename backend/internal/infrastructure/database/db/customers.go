package db

import (
	"context"
	"fmt"
	"miniapp/internal/dto"
)

func (r *repository) SaveNewCustomer(ctx context.Context, newCustomer dto.CustomerDTO) (id int, err error) {

	r.logger.Infoln("creating customer")

	query := `
			INSERT INTO customers (
				tg_name,
				tg_id,
				phone_number
			)
			VALUES ($1, $2, $3)
			RETURNING id
			`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, newCustomer.Name, newCustomer.TgId, newCustomer.PhoneNumber).Scan(&id)
	if id == 0 {
		return id, fmt.Errorf("failed to create customer")
	}

	r.logger.Infoln("created customer with id:", id)

	return id, nil
}

func (r *repository) LoadCustomer(ctx context.Context, tg_id int) (customer dto.CustomerDTO, err error) {

	r.logger.Infoln("loading customer with tg_id:", tg_id)

	query := ` 
		SELECT 
			id, 
			tg_name, 
			tg_id, 
			phone_number, 
			created_at
		FROM customers 
		WHERE tg_id = $1
		`

	r.logger.Traceln("SQL Query:", formatQuery(query))
	row := r.client.QueryRow(ctx, query, tg_id)
	err = row.Scan(&customer.Id, &customer.Name, &customer.TgId, &customer.PhoneNumber, &customer.CreatedAt)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *repository) UpdateCustomer(ctx context.Context, customer dto.CustomerDTO) (id int, err error) {

	r.logger.Infoln("updating customer, id:", customer.Id)

	query := `
			UPDATE customers
			SET
				tg_name = $2,
				phone_number = $3

			WHERER id = $1
			VALUES ($1, $2, $3)
			RETURNING id
			`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, customer.Id, customer.Name, customer.PhoneNumber).Scan(&id)
	if id == 0 {
		return id, fmt.Errorf("failed to update customer")
	}

	r.logger.Infoln("successful updated customer, id:", id)

	return id, nil
}
