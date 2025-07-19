package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/errors"
	"AuthInGo/services"
	"AuthInGo/utils"

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

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Call the service to create user
	err := uc.UserService.CreateUser()
	if err != nil {
		appErr := errors.NewAppError("Failed to create user", http.StatusBadRequest, err)
		errors.WriteError(w, appErr)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User created successfully", map[string]string{"status": "registered"})
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var payload dto.LoginUserRequestDTO

	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		appErr := errors.NewAppError("Something went wrong while logging in", http.StatusBadRequest, jsonErr)
		errors.WriteError(w, appErr)
		return
	}

	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", validationErr)
	}

	// Call the service to login user
	jwtToken, err := uc.UserService.LoginUser(&payload)
	if err != nil {
		status := http.StatusUnauthorized
		if err.Error() == "No user found with email: "+payload.Email {
			status = http.StatusNotFound
			appErr := errors.NewAppError("User not found with the provided email", status, err)
			errors.WriteError(w, appErr)
			return
		}
		appErr := errors.NewAppError(err.Error(), status, err)
		errors.WriteError(w, appErr)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)
}
