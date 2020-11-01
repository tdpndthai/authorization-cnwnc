package repository

import (
	"admin-go/models"
)

type IUserRepo interface {
	Login(loginModel *models.LoginModel) *models.UserModel
	GetAll() []models.UserModel
	GetAllPermissionOfUser(userId int) []models.PermissionUserModel
	GetById(id int) *models.UserModel
	GetAllRoleOfUser(userId int) []models.RoleUserModel
	GetAllRole() []models.Role
	GetAllPermission() []models.PermissionModel
	UpdateUser(user *models.UpdateUserModel) error
	CreateUser(user *models.CreateUserModel) error
	DeleteUser(userId int) error
}
