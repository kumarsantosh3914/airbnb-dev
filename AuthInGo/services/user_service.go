package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type UserService interface {
	CreateUser() error
	GetByID() error
	LoginUser() (string, error)
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

func (u *UserServiceImpl) LoginUser() (string, error) {
	email := "santosh@gmail.com"
	password := "Santosh@123323"

	user, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		fmt.Println("Error fetching user by email")
		return "", err
	}

	if user == nil {
		fmt.Println("No user found with the given email")
		return "", fmt.Errorf("No user found with email: %s", email)
	}

	isPasswordValid := utils.CheckPasswordHash(password, user.Password)
	if !isPasswordValid {
		fmt.Println("Password does not match")
		return "", nil
	}

	payload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}

	fmt.Println("JWT Token:", tokenString)

	return tokenString, nil
}
