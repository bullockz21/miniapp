package dto

type CustomerDTO struct {
	CustomerId      int      `json:"customer_id"`
	CustomerTgId    string   `json:"customer_telegram_id"`
	CustomerName    string   `json:"customer_name"`
	CustomerAddress []string `json:"customer_address"`
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
