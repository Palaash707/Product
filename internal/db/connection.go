package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() error {
    var err error
    DB, err = sql.Open("postgres", "user=postgres dbname=postgres password=something sslmode=disable")
    if err != nil {
        return err
    }
    return DB.Ping()
}
