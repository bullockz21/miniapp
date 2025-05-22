package service

import (
	"miniapp/internal/dto"
	"miniapp/internal/infrastructure/database/storage"
)

// Customer Service

type CustomerService interface {
	CheckAuthorization()
	Authorization(cust dto.CustomerDTO) string
	GetCustomer()
	AddCustomer()
}

type CustomerServiceImpl struct {
	Repo storage.Storage
}

// Menu Service

type MenuService interface {
}

type MenuServiceImpl struct {
	Repo storage.Storage
}

// Order Service

type OrderService interface {
}

type OrderServiceImpl struct {
	Repo storage.Storage
}

// Bucket Service

type BucketService interface {
}

type BucketServiceImpl struct {
	Repo storage.Storage
}
