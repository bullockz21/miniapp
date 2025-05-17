package dto

import (
	modules_customer "miniapp/internal/domain/customer"
	modules_menu "miniapp/internal/domain/menu"
	modules_order "miniapp/internal/domain/order"
)

// Mapping to DTO structure

func FromCustomerToDTO(u modules_customer.Customer) CustomerDTO {
	return CustomerDTO{
		CustomerId:      u.CustomerId,
		CustomerTgId:    u.CustomerTgId,
		CustomerName:    u.CustomerName,
		CustomerAddress: u.CustomerAddress,
	}
}

func FromMenuToDTO(m modules_menu.Menu) MenuDTO {
	return MenuDTO{
		MenuId:       m.MenuId,
		MenuName:     m.MenuName,
		MenuCategory: m.MenuCategory,
		MenuPrice:    m.MenuPrice,
	}
}

func FromOrderToDTO(o modules_order.Order) OrderDTO {
	return OrderDTO{
		OrderId:     o.OrderId,
		CustomerId:  o.OrderCustomerId,
		MenuId:      o.OrderMenuId,
		OrderStatus: o.OrderStatus,
		Address:     o.OrderAddress,
	}
}

// Mapping from DTO structure

func FromDTOToCustomer(u CustomerDTO) modules_customer.Customer {
	return modules_customer.Customer{
		CustomerId:      u.CustomerId,
		CustomerTgId:    u.CustomerTgId,
		CustomerName:    u.CustomerName,
		CustomerAddress: u.CustomerAddress,
	}
}

func FromDTOToMenu(m MenuDTO) modules_menu.Menu {
	return modules_menu.Menu{
		MenuId:       m.MenuId,
		MenuName:     m.MenuName,
		MenuCategory: m.MenuCategory,
		MenuPrice:    m.MenuPrice,
	}
}

func FromDTOToOrder(o OrderDTO) modules_order.Order {
	return modules_order.Order{
		OrderId:         o.OrderId,
		OrderCustomerId: o.CustomerId,
		OrderMenuId:     o.MenuId,
		OrderStatus:     o.OrderStatus,
		OrderAddress:    o.Address,
	}
}
