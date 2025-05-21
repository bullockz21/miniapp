package repository

import "sync"

func NewRepo() Repository {
	return &RepositoryImpl{
		Storage: sync.Map{},
	}
}

func (r *RepositoryImpl) GetCustomer() {

}
