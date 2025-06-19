package customer

import "time"

type Customer struct {
	Id          int       `json:"customer_id"`
	TgId        string    `json:"customer_telegram_id"`
	Name        string    `json:"customer_name"`
	Address     []string  `json:"customer_address"`
	CreatedAt   time.Time `json:"customer_created_at"`
	PhoneNumber string    `json:"phone_number"`
}
