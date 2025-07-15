package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
)

type UserService interface {
	CreateUser() error
	GetByID() error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetByID() error {
	fmt.Println("Creating user in userService")
	u.userRepository.GetByID()
	return nil
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user in UserService")
	username := "santosh12"
	email := "santosh@gmail.com"
	password := "Santosh@123323"

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	u.userRepository.Create(
		username,
		email,
		hashedPassword,
	)
	return nil
}
