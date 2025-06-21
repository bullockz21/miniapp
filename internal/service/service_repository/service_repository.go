package service_repository

import (
	"miniapp/internal/infrastructure/database/storage"
	"miniapp/internal/service"
	"miniapp/internal/service/customer/customer_repository"
)

type Service_repo struct {
	CustomerRepo customer_repository.Customer
}

func NewServiceRepository(s storage.Storage) Service_repo {
	return Service_repo{
		CustomerRepo: service.NewCustomerService(s),
	}
}
