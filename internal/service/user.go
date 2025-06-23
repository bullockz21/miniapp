package service

import (
	"context"
	"miniapp/internal/domain/user"
	"miniapp/internal/dto"
	"miniapp/internal/infrastructure/database/storage"
	"miniapp/internal/mappers"
	"miniapp/internal/service/user_repository"
)

type User struct {
	storage storage.Storage
}

func NewUserService(s storage.Storage) user_repository.User {
	return &User{
		storage: s,
	}
}

func (u *User) AuthUser(ctx context.Context, userAuth dto.AuthDTO) (token string, err error) {
	if userAuth.Login == "admin" && userAuth.Password == "admin" {
		token, err = CreateToken()
	}
	return token, err
}

func (u *User) Create(ctx context.Context, webUser dto.WebCreateUserDTO) (id int, err error) {
	NewUser := mappers.FromWebCreateUserToUser(webUser)
	NewUser.PasswordHash, err = user.CreatePasswordHash(NewUser.Password)
	if err != nil {
		return id, err
	}
	DBuser := mappers.FromUserToDBCreateUser(NewUser)
	id, err = u.storage.CreateNewUser(ctx, DBuser)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (u *User) Load(ctx context.Context, id int) (webCustomer dto.WebLoadCustomerDTO, err error) {
	return
}
func (u *User) Update(ctx context.Context, wc dto.WebUpdateCustomerDTO) (id int, err error) {
	return
}
func (u *User) LoadList(ctx context.Context) (webCustomers []dto.WebLoadCustomerDTO, err error) {
	return
}
