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

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbUri := fmt.Sprintf("host=db user=%s dbname=%s sslmode=disable password=%s", username, dbName, password) //Build connection string
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
