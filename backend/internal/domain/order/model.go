package modules_order

type Order interface {
	GetOrderId() int
	GetOrderCustomerId() int
	GetOrderMenuId() int
	GetOrderStatus() int
	GetOrderAddress() string
	SetOrderId(id int)
	SetOrderCustomerId(id int)
	SetOrderMenuId(id int)
	SetOrderStatus(stat int)
	SetOrderAddress(addr string)
}

type OrderImpl struct {
	OrderId         int    `json:"order_id"`
	OrderCustomerId int    `json:"customer_id"`
	OrderMenuId     int    `json:"menu_id"`
	OrderStatus     int    `json:"order_status"`
	OrderAddress    string `json:"address"`
}
