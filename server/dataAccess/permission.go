package dataAccess

type Permission struct {
	PermissionId           int    `gorm:"primary_key;auto_increment" json:"permissionId"`
	PermissionTitle        string `json:"permissionTitle"`
	PermissionConstantName string `json:"permissionConstantName"`
}
