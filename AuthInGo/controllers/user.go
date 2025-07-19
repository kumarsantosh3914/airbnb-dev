package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"strconv"

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

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching user by ID in UserController")
	// Extract userid from url parameters
	userId := r.URL.Query().Get("id")
	if userId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User ID is required", fmt.Errorf("Missing user ID"))
		return
	}

	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid user ID format", err)
		return
	}

	user, err := uc.UserService.GetByID(id)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user", err)
		return
	}

	if user == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "User not found", fmt.Errorf("No user found with ID: %d", id))
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User fetched successfully", user)
	fmt.Println("User fetched successfully:", user)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value("payload").(dto.CreateUserRequestDTO)

	fmt.Println("Payload received:", payload)

	user, err := uc.UserService.CreateUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User created successfully", user)

	fmt.Println("User created successfully:", user)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value("payload").(dto.LoginUserRequestDTO)

	// Call the service to login user
	jwtToken, err := uc.UserService.LoginUser(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusUnauthorized, "Failed to login user", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)
}
