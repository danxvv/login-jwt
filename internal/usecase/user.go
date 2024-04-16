package usecase

import (
	"fmt"
	"login-user/internal/entity"
)

type UserCaseUse struct {
	userRepo UserRepo
}

func NewUserUsecase(u UserRepo) *UserCaseUse {
	return &UserCaseUse{userRepo: u}
}

func (u *UserCaseUse) RegisterUser(user entity.User) error {
	exist, err := u.userRepo.GetUserByUsername(user.Username)
	if err == nil && exist.Username == user.Username {
		return fmt.Errorf("username %s already exist", user.Username)
	}
	exist, err = u.userRepo.GetUserByEmail(user.Email)
	if err == nil && exist.Email == user.Email {
		return fmt.Errorf("email %s already exist", user.Email)
	}
	exist, err = u.userRepo.GetUserByPhone(user.Phone)
	if err == nil && exist.Phone == user.Phone {
		return fmt.Errorf("phone %s already exist", user.Phone)
	}
	err = user.HashPassword(user.Password)
	if err != nil {
		return err
	}
	return u.userRepo.CreateUser(user)
}

func (u *UserCaseUse) LoginUser(username, password string) (entity.User, error) {
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		user, err = u.userRepo.GetUserByEmail(username)
		if err != nil {
			return user, err
		}
	}
	if !user.ComparePassword(password) {
		return user, fmt.Errorf("invalid password")
	}
	return user, nil
}

func (u *UserCaseUse) GetUserByID(id string) (entity.User, error) {
	return u.userRepo.GetUserByID(id)
}
