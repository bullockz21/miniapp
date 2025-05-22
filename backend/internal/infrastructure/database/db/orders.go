package db

import (
	"context"
	"fmt"
	"miniapp/internal/dto"
)

func (r *repository) SaveNewOrder(ctx context.Context, newOrder dto.OrderDTO) (id int, err error) {

	r.logger.Infoln("saving new order")

	query := `
			INSERT INTO orders (
				customer_id,
				delivery_type_id,
				menu_id,
				order_number
			)
			VALUES ($1, $2, $3, $4)
			RETURNING id
			`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, newOrder.CustomerId, newOrder.DeliveryType, newOrder.MenuId, newOrder.Number).Scan(&id)
	if id == 0 {
		return id, fmt.Errorf("failed to create user")
	}

	r.logger.Infoln("created new order with id:", id)

	return id, nil
}

func (r *repository) LoadOrdersByOrderNum(ctx context.Context, order_num int) (orders []dto.OrderDTO, err error) {

	r.logger.Infoln("loading orders by order num:", order_num)

	query := ` 
		SELECT 
			id,
			customer_id,
			delivery_type_id,
			menu_id,
			order_number
			order_state,
			order_date
		FROM customers 
		WHERE id = $1
		`

	r.logger.Traceln("SQL Query:", formatQuery(query))
	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tempOrder := dto.OrderDTO{}
		err = rows.Scan(&tempOrder.Id, &tempOrder.CustomerId, &tempOrder.DeliveryType, &tempOrder.MenuId, &tempOrder.Number, &tempOrder.Status, &tempOrder.Date)

		if err != nil {
			return nil, err
		}

		orders = append(orders, tempOrder)
	}

	return orders, nil
}

func (r *repository) UpdateOrder(ctx context.Context, order dto.OrderDTO) (id int, err error) {

	r.logger.Infoln("updating customer, id:", order.Id)

	query := `
			UPDATE orders
			SET
				customer_id,
				delivery_type_id,
				menu_id,
				order_number
				order_state,
				order_date
			WHERE id = $1
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING id
			`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, order.Id, order.CustomerId, order.DeliveryType, order.MenuId, order.Number, order.Status, order.Date).Scan(&id)
	if id == 0 {
		return id, fmt.Errorf("failed to update order")
	}

	r.logger.Infoln("successful updated order, id:", id)

	return id, nil
}
