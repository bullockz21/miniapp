package service

import (
	"miniapp/internal/dto"
	"miniapp/internal/infrastructure/database/storage"
)

func NewCustomerService(repo storage.Storage) CustomerService {
	return &CustomerServiceImpl{Repo: repo}
}

func (u *CustomerServiceImpl) CheckAuthorization() {

}

func (u *CustomerServiceImpl) Authorization(cust dto.CustomerDTO) string {
	return "token"
}

func (u *CustomerServiceImpl) GetCustomer() {

}

func (u *CustomerServiceImpl) AddCustomer() {

}
