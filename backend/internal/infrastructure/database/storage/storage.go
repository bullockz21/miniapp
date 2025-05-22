package storage

import (
	"context"
	"miniapp/internal/dto"
)

type Storage interface {
	SaveCustomer(ctx context.Context, newCustomer dto.CustomerDTO) (id int, err error)
	LoadCustomer(ctx context.Context, tg_id int) (dto.CustomerDTO, error)
	UpdateCustomer(ctx context.Context, customer dto.CustomerDTO) (id int, err error)

	SaveUser(ctx context.Context, newUser dto.UserDTO) (id int, err error)
	LoadUser(ctx context.Context, id int) (user dto.UserDTO, err error)
	UpdateUser(ctx context.Context, user dto.UserDTO) (id int, err error)
	DeleteUser(ctx context.Context, id int) (err error)
}
