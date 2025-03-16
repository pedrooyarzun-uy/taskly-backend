package db

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var DB *sqlx.DB

func Init() {
	url := os.Getenv("DB_URL")
	db_name := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")

	DB, err := sqlx.Open("postgres", "postgres://"+username+":"+password+"@"+url+":"+port+"/"+db_name)

	if err != nil {
		log.Fatal("String connection failed: ", err)
	}

	if DB.Ping() != nil {
		log.Fatal("Ping didn't work: ", err)
	}
}

func Close() {
	if err := DB.Close(); err != nil {
		log.Fatal("Close failed: ", err)
	}
}
