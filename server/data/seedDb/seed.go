package seeddb

import (
	"admin-go/dataAccess"
	"log"

	"github.com/jinzhu/gorm"
)

func InitializerDb(db *gorm.DB) {
	InitUser(db)
	InitRole(db)
	InitPermission(db)
	InitUserRole(db)
	InitPermission(db)
	InitRolePermission(db)
	InitCategory(db)
}

func InitUser(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&dataAccess.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}

func InitRole(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&dataAccess.Role{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}
func InitPermission(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&dataAccess.Permission{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}
func InitUserRole(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&dataAccess.UserRole{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}

func InitRolePermission(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&dataAccess.RolePermission{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}

func InitCategory(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&dataAccess.Category{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}
