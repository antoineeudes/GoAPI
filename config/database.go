package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "db"
	port     = 5432
	user     = "username"
	password = "secret"
	dbname   = "dev"
)

func DatabaseInit() {
	var err error
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", dbinfo)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
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

func Db() *sql.DB {
	return db
}
