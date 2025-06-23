package service_repository

import (
	"miniapp/internal/infrastructure/database/storage"
	"miniapp/internal/service"
	"miniapp/internal/service/customer_repository"
	"miniapp/internal/service/user_repository"
)

type Service_repo struct {
	CustomerRepo customer_repository.Customer
	UserRepo     user_repository.User
}

func NewServiceRepository(s storage.Storage) Service_repo {
	return Service_repo{
		CustomerRepo: service.NewCustomerService(s),
		UserRepo:     service.NewUserService(s),
	}
}
