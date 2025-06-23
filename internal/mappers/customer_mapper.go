package mappers

import (
	"miniapp/internal/domain/customer"
	"miniapp/internal/dto"
)

func FromWebCreateCustomerToCustomer(wc dto.WebCreateCustomerDTO) customer.Customer {
	return customer.Customer{
		TgId: wc.TgId,
		Name: wc.Name,
	}
}

func FromWebUpdateCustomerToCustomer(wc dto.WebUpdateCustomerDTO) customer.Customer {
	return customer.Customer{
		Id:          wc.Id,
		Name:        wc.Name,
		PhoneNumber: wc.PhoneNumber,
	}
}

func FromCustomerToDBCreateCustomer(c customer.Customer) dto.DBCreateCustomerDTO {
	return dto.DBCreateCustomerDTO{
		TgId: c.TgId,
		Name: c.Name,
	}
}

func FromCustomerToDBLoadCustomer(c customer.Customer) dto.DBLoadCustomerDTO {
	return dto.DBLoadCustomerDTO{
		TgId: c.TgId,
		Name: c.Name,
	}
}

func FromCustomerToDBUpdateCustomer(c customer.Customer) dto.DBUpdateCustomerDTO {
	return dto.DBUpdateCustomerDTO{
		Id:          c.Id,
		Name:        c.Name,
		PhoneNumber: c.PhoneNumber,
	}
}

func FromCustomerToWebLoadCustomer(c customer.Customer) dto.WebLoadCustomerDTO {
	return dto.WebLoadCustomerDTO{
		Id:          c.Id,
		TgId:        c.TgId,
		Name:        c.Name,
		PhoneNumber: c.PhoneNumber,
		CreatedAt:   c.CreatedAt,
	}
}

func FromDBLoadCustomerToCustomer(dc dto.DBLoadCustomerDTO) customer.Customer {
	return customer.Customer{
		Id:          dc.Id,
		TgId:        dc.TgId,
		Name:        dc.Name,
		PhoneNumber: dc.PhoneNumber,
		CreatedAt:   dc.CreatedAt,
	}
}
