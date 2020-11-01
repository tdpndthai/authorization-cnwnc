package dataAccess

type Role struct {
	RoleId    int    `gorm:"primary_key;auto_increment" json:"roleId"`
	RoleTitle string `json:"roleTitle"`
}
