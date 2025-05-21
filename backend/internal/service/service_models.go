package service

import (
	"miniapp/internal/dto"
	"miniapp/internal/infrastructure/database/repository"
)

// Customer Service

type CustomerService interface {
	CheckAuthorization()
	Authorization(cust dto.CustomerDTO) string
	GetCustomer()
	AddCustomer()
}

type CustomerServiceImpl struct {
	Repo repository.Repository
}

// Menu Service

type MenuService interface {
}

type MenuServiceImpl struct {
	Repo repository.Repository
}

// Order Service

type OrderService interface {
}

type OrderServiceImpl struct {
	Repo repository.Repository
}

// Bucket Service

type BucketService interface {
}

type BucketServiceImpl struct {
	Repo repository.Repository
}
