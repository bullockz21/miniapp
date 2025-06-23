package dto

import "time"

// id SERIAL PRIMARY KEY,
// type_id INT NOT NULL,
// name VARCHAR(50) UNIQUE NOT NULL,
// price NUMERIC,
// description VARCHAR(1000)

type MenuDTO struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Category    int     `json:"type_id"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// id SERIAL PRIMARY KEY,
// customer_id INT NOT NULL,
// delivery_type_id INT NOT NULL,
// menu_id INT NOT NULL,
// order_number INT NOT NULL,
// order_state_id INT DEFAULT 1,
// order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP

type OrderDTO struct {
	Id           int       `json:"id"`
	CustomerId   int       `json:"customer_id"`
	MenuId       int       `json:"menu_id"`
	DeliveryType int       `json:"delivery_type_id"`
	Number       int       `json:"order_num"`
	Status       int       `json:"order_state_id"`
	Date         time.Time `json:"order_date"`
}
