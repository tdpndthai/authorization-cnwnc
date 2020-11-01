package models

type PermissionModel struct {
	PermissionId           int    `json:"permissionId"`
	PermissionTitle        string `json:"permissionTitle"`
	PermissionConstantName string `json:"permissionConstantName"`
}
