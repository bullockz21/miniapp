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

type MenuDTO struct {
	MenuId       int    `json:"menu_id"`
	MenuName     string `json:"menu_name"`
	MenuCategory int    `json:"menu_category"`
	MenuPrice    int    `json:"menu_price"`
}

type OrderDTO struct {
	OrderId     int    `json:"order_id"`
	CustomerId  int    `json:"customer_id"`
	MenuId      int    `json:"menu_id"`
	OrderStatus int    `json:"order_status"`
	Address     string `json:"address"`
}
