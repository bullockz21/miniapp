package modules_order

type Order interface {
}

type OrderImpl struct {
	OrderId         int    `json:"order_id"`
	OrderCustomerId int    `json:"customer_id"`
	OrderMenuId     int    `json:"menu_id"`
	OrderStatus     int    `json:"order_status"`
	OrderAddress    string `json:"address"`
}
