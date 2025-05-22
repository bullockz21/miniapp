package dto

type CustomerDTO struct {
	Id          int      `json:"customer_id"`
	TgId        string   `json:"customer_telegram_id"`
	Name        string   `json:"customer_name"`
	Address     []string `json:"customer_address"`
	PhoneNumber string   `json:"phone_number"`
	CreatedAt   string   `json:"created_at"`
}

type UserDTO struct {
	Id           int    `json:"id"`
	UserName     string `json:"user_name"`
	RoleId       int    `json:"role_id"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
	CreatedAt    string `json:"created_at"`
	DeletedAt    string `json:"deleted_at"`
}

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
	Id           int    `json:"id"`
	CustomerId   int    `json:"customer_id"`
	MenuId       int    `json:"menu_id"`
	DeliveryType int    `json:"delivery_type_id"`
	Number       int    `json:"order_num"`
	Status       int    `json:"order_state_id"`
	Date         string `json:"order_date"`
}
