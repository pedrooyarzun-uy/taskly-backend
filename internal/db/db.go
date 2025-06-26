package db

import (
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

		dsn := os.Getenv("PSQL_URL")

		var err error
		DB, err = sqlx.Open("postgres", dsn)

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
