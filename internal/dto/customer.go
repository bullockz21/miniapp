package dto

import "time"

// id SERIAL PRIMARY KEY,
// tg_name VARCHAR(100) UNIQUE NOT NULL,
// tg_id INT UNIQUE NOT NULL,
// phone_number VARCHAR(50),
// created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

type WebCreateCustomerDTO struct {
	TgId int    `json:"customer_telegram_id"`
	Name string `json:"customer_name"`
}

type WebLoadCustomerDTO struct {
	Id          int       `json:"customer_id"`
	TgId        int       `json:"customer_telegram_id"`
	Name        string    `json:"customer_name"`
	PhoneNumber string    `json:"customer_phone_number"`
	CreatedAt   time.Time `json:"customer_created_at"`
}

type WebUpdateCustomerDTO struct {
	Id          int    `json:"customer_id"`
	Name        string `json:"customer_name"`
	PhoneNumber string `json:"customer_phone_number"`
}

type DBCreateCustomerDTO struct {
	TgId int    `json:"customer_telegram_id"`
	Name string `json:"customer_name"`
}

type DBUpdateCustomerDTO struct {
	Id          int    `json:"customer_id"`
	Name        string `json:"customer_name"`
	PhoneNumber string `json:"customer_phone_number"`
}

type DBLoadCustomerDTO struct {
	Id          int       `json:"customer_id"`
	TgId        int       `json:"customer_telegram_id"`
	Name        string    `json:"customer_name"`
	PhoneNumber string    `json:"customer_phone_number"`
	CreatedAt   time.Time `json:"customer_created_at"`
}
