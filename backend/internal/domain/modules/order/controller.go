package modules_order

func NewOrder() Order {
	return Order{}
}

// Getters
func (o Order) GetOrderId() int {
	return o.OrderId
}

func (o Order) GetOrderUserId() int {
	return o.OrderUserId
}

func (o Order) GetOrderMenuId() int {
	return o.OrderMenuId
}

func (o Order) GetOrderStatus() int {
	return o.OrderStatus
}

func (o Order) GetOrderAddress() string {
	return o.OrderAddress
}

func (o *Order) SetOrderId(address string) {
	o.OrderAddress = address
}
