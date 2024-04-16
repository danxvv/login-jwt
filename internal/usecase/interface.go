package usecase

import "login-user/internal/entity"

type (
	UserRepo interface {
		CreateUser(user entity.User) error
		GetUserByUsername(username string) (entity.User, error)
		GetUserByEmail(email string) (entity.User, error)
		GetUserByID(id string) (entity.User, error)
		GetUserByPhone(phone string) (entity.User, error)
	}

	User interface {
		RegisterUser(user entity.User) error
		LoginUser(username, password string) (entity.User, error)
		GetUserByID(id string) (entity.User, error)
	}
)
