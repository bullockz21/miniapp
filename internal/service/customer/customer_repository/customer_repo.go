package customer_repository

import "miniapp/internal/dto"

type Customer interface {
	Create(webCustomer dto.WebCreateCustomerDTO) (id int, err error)
	Load(id int) (webCustomer dto.WebLoadCustomerDTO, err error)
	Update(wc dto.WebUpdateCustomerDTO) (id int, err error)
	LoadList() (webCustomers []dto.WebLoadCustomerDTO, err error)
}
