package storage

import (
	"context"
	"miniapp/internal/dto"
)

type Storage interface {
	SaveNewCustomer(ctx context.Context, newCustomer dto.CustomerDTO) (id int, err error)
	LoadCustomer(ctx context.Context, tg_id int) (dto.CustomerDTO, error)
	UpdateCustomer(ctx context.Context, customer dto.CustomerDTO) (id int, err error)

	SaveNewUser(ctx context.Context, newUser dto.UserDTO) (id int, err error)
	LoadUser(ctx context.Context, id int) (user dto.UserDTO, err error)
	UpdateUser(ctx context.Context, user dto.UserDTO) (id int, err error)
	DeleteUser(ctx context.Context, id int) (err error)

	UpdateOrder(ctx context.Context, order dto.OrderDTO) (id int, err error)
	LoadOrdersByOrderNum(ctx context.Context, order_num int) (orders []dto.OrderDTO, err error)
	SaveNewOrder(ctx context.Context, newOrder dto.OrderDTO) (id int, err error)
}
