package dto

import "time"

type CustomerDTO struct {
	Id          int       `json:"customer_id"`
	TgId        string    `json:"customer_telegram_id"`
	Name        string    `json:"customer_name"`
	Address     []string  `json:"customer_address"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type WebUpdateCustomerDTO struct {
	TgId        string `json:"customer_telegram_id"`
	PhoneNumber string `json:"phone_number"`
}

type WebCreateMenuDTO struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Category    int     `json:"type_id"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
