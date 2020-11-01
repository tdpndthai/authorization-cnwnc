package models

type RolePermissions struct {
	RolePermissionId int `json:"rolePermissionId"`
	RoleId           int `json:"roleId"`
	PermissionId     int `json:"permissionId"`
}
