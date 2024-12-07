package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// ConnectDB initializes the database connection
func ConnectDB() {
    var err error
    // Update the connection string with your credentials
    DB, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=something dbname=productdb sslmode=disable")
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatalf("Cannot reach database: %v", err)
    }

    log.Println("Connected to database successfully")
}
