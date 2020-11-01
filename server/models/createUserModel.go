package models

type CreateUserModel struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Email    string `json:"email"`
	RoleId   int `json:"roleId"`
}
