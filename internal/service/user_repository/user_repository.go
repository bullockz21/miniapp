package user_repository

import (
	"context"
	"miniapp/internal/dto"
)

type User interface {
	AuthUser(ctx context.Context, userAuth dto.AuthDTO) (token string, err error)
	Create(ctx context.Context, webUser dto.WebCreateUserDTO) (id int, err error)
	Load(ctx context.Context, id int) (webCustomer dto.WebLoadCustomerDTO, err error)
	Update(ctx context.Context, wc dto.WebUpdateCustomerDTO) (id int, err error)
	LoadList(ctx context.Context) (webCustomers []dto.WebLoadCustomerDTO, err error)
}
