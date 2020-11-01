package models

type UpdateUserModel struct {
	UserId           int                   `json:"userId"`
	UserName         string                `json:"username"`
	PassWord         string                `json:"password"`
	Email            string                `json:"email"`
	DeletePermission []PermissionUserModel `json:"deletePermission"`
	DeleteRole       []RoleUserModel       `json:"deleteRole"`
	RoleUser         []RoleUserModel       `json:"roleUser"`
}
