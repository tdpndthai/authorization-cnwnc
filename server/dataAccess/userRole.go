package dataAccess

type UserRole struct {
	UserRoleId int `gorm:"primary_key;auto_increment" json:"userRoleId"`
	UserId     int `json:"userId"`
	RoleId     int `json:"roleId"`
}
