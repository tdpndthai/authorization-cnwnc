package models

type UserModel struct {
	UserId       int    `json:"userId"`
	UserName     string `json:"username"`
	PassWord     string `json:"password"`
	Email        string `json:"email"`
	IsSuperAdmin bool   `json:"isSuperAdmin"`
}
