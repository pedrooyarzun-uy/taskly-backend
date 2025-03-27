package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
)

var DB *sqlx.DB
var once sync.Once

func Init() {
	once.Do(func() {
		host := os.Getenv("DB_URL")
		dbName := os.Getenv("DB_NAME")
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		port := os.Getenv("DB_PORT")

		conString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, username, password, dbName,
		)

		var err error
		DB, err = sqlx.Open("postgres", conString)

		if err != nil {
			log.Fatal("String connection failed: ", err)
		}

		if err = DB.Ping(); err != nil {
			log.Fatal("Ping didn't work: ", err)
		}

	})
}

func Close() {
	if err := DB.Close(); err != nil {
		log.Fatal("Close failed: ", err)
	}
}
