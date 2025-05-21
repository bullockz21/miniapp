package modules_order

func NewOrder() Order {
	return &OrderImpl{}
}

// Getters

func (o OrderImpl) GetOrderId() int {
	return o.OrderId
}

func (o OrderImpl) GetOrderCustomerId() int {
	return o.OrderCustomerId
}

func (o OrderImpl) GetOrderMenuId() int {
	return o.OrderMenuId
}

func (o OrderImpl) GetOrderStatus() int {
	return o.OrderStatus
}

func (o OrderImpl) GetOrderAddress() string {
	return o.OrderAddress
}

// Setters

func (o *OrderImpl) SetOrderId(id int) {
	o.OrderId = id
}

func (o *OrderImpl) SetOrderCustomerId(id int) {
	o.OrderCustomerId = id
}

func (o *OrderImpl) SetOrderMenuId(id int) {
	o.OrderMenuId = id
}

func (o *OrderImpl) SetOrderStatus(stat int) {
	o.OrderStatus = stat
}

func (o *OrderImpl) SetOrderAddress(addr string) {
	o.OrderAddress = addr
}
