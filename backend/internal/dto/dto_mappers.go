package dto

import (
	modules_menu "miniapp/internal/domain/modules/menu"
	modules_order "miniapp/internal/domain/modules/order"
	modules_user "miniapp/internal/domain/modules/user"
)

// Mapping to DTO structure

func FromUserToDTO(u modules_user.User) UserDTO {
	return UserDTO{
		UserId:      u.UserId,
		UserTgId:    u.UserTgId,
		UserName:    u.UserName,
		UserAddress: u.UserAddress,
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
		UserId:      o.OrderUserId,
		MenuId:      o.OrderMenuId,
		OrderStatus: o.OrderStatus,
		Address:     o.OrderAddress,
	}
}

// Mapping from DTO structure

func FromDTOToUser(u UserDTO) modules_user.User {
	return modules_user.User{
		UserId:      u.UserId,
		UserTgId:    u.UserTgId,
		UserName:    u.UserName,
		UserAddress: u.UserAddress,
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
		OrderId:      o.OrderId,
		OrderUserId:  o.UserId,
		OrderMenuId:  o.MenuId,
		OrderStatus:  o.OrderStatus,
		OrderAddress: o.Address,
	}
}
