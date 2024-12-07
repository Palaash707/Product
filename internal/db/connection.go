package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// ConnectDB establishes the database connection
func ConnectDB() {
    var err error
    DB, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=something dbname=productdb sslmode=disable")
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    // Check if the connection works
    if err = DB.Ping(); err != nil {
        log.Fatalf("Cannot reach database: %v", err)
    }

    log.Println("Connected to database successfully")
}
