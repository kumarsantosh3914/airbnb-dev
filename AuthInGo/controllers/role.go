package controllers

import (
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoleController struct {
	roleService services.RoleService
}

func NewRoleController(roleService services.RoleService) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (rc *RoleController) GetRoleById(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id") // Extract the roleId from the URL
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", nil)
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", nil)
		return
	}

	role, err := rc.roleService.GetRoleById(id)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to get role", nil)
		return
	}

	if role == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "Role not found", fmt.Errorf("role with ID %d not found", roleId))
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role fetched successfully", role)
}

func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := rc.roleService.GetAllRoles()
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to get roles", nil)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Roles fetched successfully", roles)
}
