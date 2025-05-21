package repository

import (
	"sync"
)

type Repository interface {
	GetCustomer()
}

type RepositoryImpl struct {
	Storage sync.Map // Vremenno. Nujno podrubit orm ili hz chto tam
}
