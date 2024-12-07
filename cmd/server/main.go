package main

import (
	"log"

	"github.com/Palaash707/Product/internal/api"
	"github.com/Palaash707/Product/internal/db"
)

func main() {
    // Connect to the database
    err := db.ConnectDB()
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }

    // Set up and start the server
    router := api.SetupRouter()
    router.Run(":8080")
}
