package dto

import (
	modules_customer "miniapp/internal/domain/customer"
	modules_menu "miniapp/internal/domain/menu"
	modules_order "miniapp/internal/domain/order"
)

// Mapping to DTO structure

func FromCustomerToDTO(u modules_customer.Customer) CustomerDTO {
	return CustomerDTO{
		CustomerId:      u.GetCustomerId(),
		CustomerTgId:    u.GetCustomerTgId(),
		CustomerName:    u.GetCustomerName(),
		CustomerAddress: u.GetCustomerAddressList(),
	}
}

func FromMenuToDTO(m modules_menu.Menu) MenuDTO {
	return MenuDTO{
		MenuId:       m.GetMenuId(),
		MenuName:     m.GetMenuName(),
		MenuCategory: m.GetMenuCategory(),
		MenuPrice:    m.GetMenuPrice(),
	}
}

func FromOrderToDTO(o modules_order.Order) OrderDTO {
	return OrderDTO{
		OrderId:     o.GetOrderId(),
		CustomerId:  o.GetOrderCustomerId(),
		MenuId:      o.GetOrderMenuId(),
		OrderStatus: o.GetOrderStatus(),
		Address:     o.GetOrderAddress(),
	}
}

// Mapping from DTO structure

func FromDTOToCustomer(u CustomerDTO) modules_customer.Customer {
	cust := modules_customer.NewCustomer()
	cust.SetCustomerId(u.CustomerId)
	cust.SetCustomerName(u.CustomerName)
	cust.SetCustomerTgId(u.CustomerTgId)
	for _, address := range u.CustomerAddress {
		cust.AddCustomerAddress(address)
	}
	return cust
}

func FromDTOToMenu(m MenuDTO) modules_menu.Menu {
	menu := modules_menu.NewMenu()
	menu.SetMenuId(m.MenuId)
	menu.SetMenuName(m.MenuName)
	menu.SetMenuCategory(m.MenuCategory)
	menu.SetMenuPrice(m.MenuPrice)
	return menu
}

func FromDTOToOrder(o OrderDTO) modules_order.Order {
	order := modules_order.NewOrder()
	order.SetOrderId(o.OrderId)
	order.SetOrderCustomerId(o.CustomerId)
	order.SetOrderMenuId(o.MenuId)
	order.SetOrderAddress(o.Address)
	order.SetOrderStatus(o.OrderStatus)
	return order
}
