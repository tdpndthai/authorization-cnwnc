package services

import "admin-go/models"

type IUserService interface {
	Login(loginModel *models.LoginModel) map[string]interface{}
	GetAll() map[string]interface{}
	GetById(id int) map[string]interface{}
	GetDetailUser(id int) map[string]interface{}
	UpdateUser(user *models.UpdateUserModel) map[string]interface{}
	CreateUser(user *models.CreateUserModel) map[string]interface{}
	DeleteUser(userId int) map[string]interface{}
}
