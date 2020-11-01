package models

type RoleUserModel struct {
	RoleId      int                   `json:"roleId"`
	RoleTitle   string                `json:"roleTitle"`
	UserRoleId  int                   `json:"userRoleId"`
	Permissions []PermissionUserModel `json:"permissions"`
}
