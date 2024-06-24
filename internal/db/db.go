package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	connStr := os.Getenv("DATABASE_URL")

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening the databases: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connection to the database: %v", err)
	}
}
