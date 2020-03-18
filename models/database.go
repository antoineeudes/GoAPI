package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB //database

func init() {

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dbUri := fmt.Sprintf("host=db port=%d user=%s dbname=%s sslmode=disable password=%s", 5432, username, dbName, password) //Build connection string
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to database")

	db = conn
	db.Debug().AutoMigrate(&Account{}, Gallery{}, File{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
