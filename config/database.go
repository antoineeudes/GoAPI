package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// github.com/lib/pq allows the connection with postgresql server
	_ "github.com/lib/pq"
)

var db *sql.DB

// DatabaseInit realizes connection with database and create tables if they don't exist
func DatabaseInit() {
	var err error
	dbinfo := fmt.Sprintf("host=db port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err = sql.Open("postgres", dbinfo)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database!")
	createCarsTable()
}

func createCarsTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS cars(id serial,manufacturer varchar(20), design varchar(20), style varchar(20), doors int, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))")

	if err != nil {
		log.Fatal(err)
	}
}

// Db returns the database used
func Db() *sql.DB {
	return db
}
