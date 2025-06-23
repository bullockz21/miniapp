package mappers

import (
	"miniapp/internal/domain/user"
	"miniapp/internal/dto"
)

// CREATE MAPPERS

func FromWebCreateUserToUser(webUser dto.WebCreateUserDTO) user.User {
	return user.User{
		Name:     webUser.Name,
		Email:    webUser.Email,
		Password: webUser.Password,
	}
}

func FromUserToDBCreateUser(user user.User) dto.DBCreateUserDTO {
	return dto.DBCreateUserDTO{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}

func FromDBCreateUserToUser(DBUser dto.DBCreateUserDTO) user.User {
	return user.User{
		Id:    DBUser.Id,
		Name:  DBUser.Name,
		Email: DBUser.Email,
	}
}
