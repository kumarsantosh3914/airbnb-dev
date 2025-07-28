package db

import (
	"AuthInGo/models"
	"database/sql"
)

type RolePermissionRepository interface {
	GetRolePermissionById(id int64) (*models.RolePermission, error)
	GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermission, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	GetAllRolePermissions() ([]*models.RolePermission, error)
}

type RolePermissionRepositoryImpl struct {
	db *sql.DB
}

func NewRolePermissionRepository(db *sql.DB) RolePermissionRepository {
	return &RolePermissionRepositoryImpl{db: db}
}

func (r *RolePermissionRepositoryImpl) GetRolePermissionById(id int64) (*models.RolePermission, error) {
	query := "SELECT id, role_id, permission, created_at, updated_at FROM role_permissions WHERE id = ?"
	row := r.db.QueryRow(query, id)
	rolePermission := &models.RolePermission{}
	if err := row.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.Permission, &rolePermission.CreatedAt, &rolePermission.UpdatedAt); err != nil {
		return nil, err
	}
	return rolePermission, nil
}

func (r *RolePermissionRepositoryImpl) GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermission, error) {
	query := "SELECT id, role_id, permission, created_at, updated_at FROM role_permissions WHERE role_id = ?"
	rows, err := r.db.Query(query, roleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var rolePermissions []*models.RolePermission
	for rows.Next() {
		rp := &models.RolePermission{}
		if err := rows.Scan(&rp.Id, &rp.RoleId, &rp.Permission, &rp.CreatedAt, &rp.UpdatedAt); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rp)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (r *RolePermissionRepositoryImpl) AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	query := "INSERT INTO role_permissions (role_id, permission, created_at, updated_at) VALUES (?, ?, NOW(), NOW())"
	result, err := r.db.Exec(query, roleId, permissionId)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r.GetRolePermissionById(id)
}

func (r *RolePermissionRepositoryImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {
	query := "DELETE FROM role_permissions WHERE role_id = ? AND permission = ?"
	_, err := r.db.Exec(query, roleId, permissionId)
	return err
}

func (r *RolePermissionRepositoryImpl) GetAllRolePermissions() ([]*models.RolePermission, error) {
	query := "SELECT id, role_id, permission, created_at, updated_at FROM role_permissions"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var rolePermissions []*models.RolePermission
	for rows.Next() {
		rp := &models.RolePermission{}
		if err := rows.Scan(&rp.Id, &rp.RoleId, &rp.Permission, &rp.CreatedAt, &rp.UpdatedAt); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rp)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rolePermissions, nil
}
