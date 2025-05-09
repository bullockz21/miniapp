package modules_order

type Order struct {
	OrderId      int    `json:"order_id"`
	OrderUserId  int    `json:"user_id"`
	OrderMenuId  int    `json:"menu_id"`
	OrderStatus  int    `json:"order_status"`
	OrderAddress string `json:"address"`
}
