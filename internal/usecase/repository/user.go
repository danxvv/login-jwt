package repository

import (
	"login-user/internal/entity"
	"login-user/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(p *postgres.Postgres) *UserRepo {
	return &UserRepo{Postgres: p}
}

func (u *UserRepo) CreateUser(user entity.User) error {
	return u.DB.Create(&user).Error
}

func (u *UserRepo) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User
	err := u.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func (u *UserRepo) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u *UserRepo) GetUserByID(id string) (entity.User, error) {
	var user entity.User
	err := u.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func (u *UserRepo) GetUserByPhone(phone string) (entity.User, error) {
	var user entity.User
	err := u.DB.Where("phone = ?", phone).First(&user).Error
	return user, err
}
