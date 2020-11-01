package data

import (
	seeddb "admin-go/data/seedDb"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/joho/godotenv"
)

var session *gorm.DB

func GetSession() *gorm.DB {
	var err error
	DBURL := fmt.Sprintf("server=%s;user id = %s;password=%s;database=%s",
		os.Getenv("DB_SERVER_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	session, err = gorm.Open("mssql", DBURL)
	if err != nil {
		log.Fatalf("[GetSession]: %s\n", err)
	}
	return session
}

// CreateDbSession database session
func CreateDbSession() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	DBURL := fmt.Sprintf("server=%s;user id = %s;password=%s;database=%s",
		os.Getenv("DB_SERVER_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	session, err = gorm.Open("mssql", DBURL)

	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}

	seeddb.InitializerDb(session)
}
