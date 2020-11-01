package models

type PermissionUserModel struct {
	PermissionId           int    `json:"permissionId"`
	RoleId                 int    `json:"roleId"`
	RoleTitle              string `json:"roleTitle"`
	PermissionTitle        string `json:"permissionTitle"`
	PermissionConstantName string `json:"permissionConstantName"`
	RolePermissionId       int    `json:"rolePermissionId"`
}
