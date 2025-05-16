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

// Setters

func (o *Order) SetOrderId(id int) {
	o.OrderId = id
}

func (o *Order) SetOrderUserId(id int) {
	o.OrderUserId = id
}

func (o *Order) SetOrderMenuId(id int) {
	o.OrderMenuId = id
}

func (o *Order) SetOrderStatus(stat int) {
	o.OrderStatus = stat
}

func (o *Order) SetOrderAddress(addr string) {
	o.OrderAddress = addr
}
