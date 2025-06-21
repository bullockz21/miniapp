package customer_repository

import (
	"context"
	"miniapp/internal/dto"
)

type Customer interface {
	Create(cxt context.Context, webCustomer dto.WebCreateCustomerDTO) (id int, err error)
	Load(ctx context.Context, id int) (webCustomer dto.WebLoadCustomerDTO, err error)
	Update(ctx context.Context, wc dto.WebUpdateCustomerDTO) (id int, err error)
	LoadList(ctx context.Context) (webCustomers []dto.WebLoadCustomerDTO, err error)
}
