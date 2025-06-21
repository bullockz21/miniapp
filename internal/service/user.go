package service

import (
	"context"
	"miniapp/internal/dto"
	"miniapp/internal/infrastructure/database/storage"
)

type User struct {
	storage storage.Storage
}

func (u User) AuthUser(ctx context.Context, userAuth dto.AuthDTO) (token string, err error) {
	if userAuth.Login == "admin" && userAuth.Password == "admin" {
		token, err = CreateToken()
	}
	return token, err
}
