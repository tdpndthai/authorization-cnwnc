package dataAccess

type User struct {
	UserId       int    `gorm:"primary_key;auto_increment" json:"userId"`
	UserName     string `json:"username"`
	PassWord     string `json:"password"`
	Email        string `json:"email"`
	IsSuperAdmin bool   `json:"isSuperAdmin"`
}
