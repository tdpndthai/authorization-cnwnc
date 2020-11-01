package dataAccess

type RolePermission struct {
	RolePermissionId int `gorm:"primary_key;auto_increment" json:"rolePermissionId"`
	RoleId           int `json:"roleId"`
	PermissionId     int `json:"permissionId`
}
