package models

import "time"

type DetailUserModel struct {
	UserId       int       `json:"userId"`
	UserName     string    `json:"username"`
	PassWord     string    `json:"passwor"`
	Email        string    `json:"email"`
	IsSuperAdmin bool      `json:"isSuperAdmi"`
	CreatedBy    string    `json:"createdBy"`
	CreatedDate  time.Time `son:"createdDat"`
	RoleModels   [] Role
}
