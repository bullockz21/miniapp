package service_repository

import (
	"miniapp/internal/infrastructure/database/storage"
	"miniapp/internal/service/customer/customer_repository"
	"miniapp/internal/service/customer/service"
)

type Service_repo struct {
	CustomerRepo customer_repository.Customer
}

func NewServiceRepository(r storage.Storage) Service_repo {
	return Service_repo{
		CustomerRepo: service.NewCustomerService(r),
	}
}
