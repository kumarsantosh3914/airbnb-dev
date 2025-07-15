package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

type UserService interface {
	// CreateUser() error
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
	u.userRepository.Create()
	return nil
}
