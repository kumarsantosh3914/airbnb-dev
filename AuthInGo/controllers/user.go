package controllers

import (
	"AuthInGo/services"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registeruser called in UserController")
	uc.UserService.GetByID()
	w.Write([]byte("User registeration endpoint"))
}
