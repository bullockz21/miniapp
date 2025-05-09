package dto

type UserDTO struct {
	UserId      int      `json:"user_id"`
	UserTgId    string   `json:"user_telegram_id"`
	UserName    string   `json:"user_name"`
	UserAddress []string `json:"user_address"`
}

type MenuDTO struct {
	MenuId       int    `json:"menu_id"`
	MenuName     string `json:"menu_name"`
	MenuCategory int    `json:"menu_category"`
	MenuPrice    int    `json:"menu_price"`
}

type OrderDTO struct {
	OrderId     int    `json:"order_id"`
	UserId      int    `json:"user_id"`
	MenuId      int    `json:"menu_id"`
	OrderStatus int    `json:"order_status"`
	Address     string `json:"address"`
}
