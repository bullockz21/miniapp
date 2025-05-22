package dto

import (
	modules_customer "miniapp/internal/domain/customer"
)

// Mapping to DTO structure

func FromCustomerToDTO(u modules_customer.Customer) CustomerDTO {
	return CustomerDTO{
		Id:      u.GetCustomerId(),
		TgId:    u.GetCustomerTgId(),
		Name:    u.GetCustomerName(),
		Address: u.GetCustomerAddressList(),
	}
}

// func FromMenuToDTO(m modules_menu.Menu) MenuDTO {
// 	return MenuDTO{
// 		MenuId:       m.GetMenuId(),
// 		MenuName:     m.GetMenuName(),
// 		MenuCategory: m.GetMenuCategory(),
// 		MenuPrice:    m.GetMenuPrice(),
// 	}
// }

// func FromOrderToDTO(o modules_order.Order) OrderDTO {
// 	return OrderDTO{
// 		OrderId:     o.GetOrderId(),
// 		CustomerId:  o.GetOrderCustomerId(),
// 		MenuId:      o.GetOrderMenuId(),
// 		OrderStatus: o.GetOrderStatus(),
// 		Address:     o.GetOrderAddress(),
// 	}
// }

// Mapping from DTO structure

func FromDTOToCustomer(u CustomerDTO) modules_customer.Customer {
	cust := modules_customer.NewCustomer()
	cust.SetCustomerId(u.Id)
	cust.SetCustomerName(u.Name)
	cust.SetCustomerTgId(u.TgId)
	for _, address := range u.Address {
		cust.AddCustomerAddress(address)
	}
	return cust
}

// func FromDTOToMenu(m MenuDTO) modules_menu.Menu {
// 	menu := modules_menu.NewMenu()
// 	menu.SetMenuId(m.MenuId)
// 	menu.SetMenuName(m.MenuName)
// 	menu.SetMenuCategory(m.MenuCategory)
// 	menu.SetMenuPrice(m.MenuPrice)
// 	return menu
// }

// func FromDTOToOrder(o OrderDTO) modules_order.Order {
// 	order := modules_order.NewOrder()
// 	order.SetOrderId(o.OrderId)
// 	order.SetOrderCustomerId(o.CustomerId)
// 	order.SetOrderMenuId(o.MenuId)
// 	order.SetOrderAddress(o.Address)
// 	order.SetOrderStatus(o.OrderStatus)
// 	return order
// }
