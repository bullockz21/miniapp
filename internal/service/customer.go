package service

import (
	"context"
	"miniapp/internal/domain/customer"
	"miniapp/internal/dto"
	"miniapp/internal/infrastructure/database/storage"
	"miniapp/internal/mappers"
	"miniapp/internal/service/customer/customer_repository"
)

type Customer struct {
	Repository storage.Storage
}

func NewCustomerService(r storage.Storage) customer_repository.Customer {
	return &Customer{
		Repository: r,
	}
}

func (c *Customer) Create(ctx context.Context, webCustomer dto.WebCreateCustomerDTO) (id int, err error) {
	customer := mappers.FromWebCreateCustomerToCustomer(webCustomer)
	DBcustomer := mappers.FromCustomerToDBCreateCustomer(customer)
	id, err = c.Repository.CreateNewCustomer(ctx, DBcustomer)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (c *Customer) Load(ctx context.Context, id int) (webCustomer dto.WebLoadCustomerDTO, err error) {
	DBCustomer, err := c.Repository.LoadCustomer(ctx, id)
	if err != nil {
		return webCustomer, err
	}
	customer := mappers.FromDBLoadCustomerToCustomer(DBCustomer)
	webCustomer = mappers.FromCustomerToWebLoadCustomer(customer)
	return webCustomer, nil
}

func (c *Customer) LoadList(ctx context.Context) (webCustomers []dto.WebLoadCustomerDTO, err error) {
	DBcustomers, err := c.Repository.LoadCustomerList(ctx)
	if err != nil {
		return webCustomers, err
	}
	customers := make([]customer.Customer, 0)
	for _, DBcustomer := range DBcustomers {
		customers = append(customers, mappers.FromDBLoadCustomerToCustomer(DBcustomer))
	}

	webCustomers = make([]dto.WebLoadCustomerDTO, 0)
	for _, customer := range customers {
		webCustomers = append(webCustomers, mappers.FromCustomerToWebLoadCustomer(customer))
	}

	return webCustomers, nil
}

func (c *Customer) Update(ctx context.Context, wc dto.WebUpdateCustomerDTO) (id int, err error) {
	customer := mappers.FromWebUpdateCustomerToCustomer(wc)
	DBCustomer := mappers.FromCustomerToDBUpdateCustomer(customer)
	id, err = c.Repository.UpdateCustomer(ctx, DBCustomer)
	if err != nil {
		return id, err
	}
	return id, nil
}
