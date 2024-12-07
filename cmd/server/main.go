package main

import (
	"log"

	"github.com/Palaash707/Product/internal/api"
	"github.com/Palaash707/Product/internal/db"
)

func main() {
    // Initialize the database connection
    db.ConnectDB()

    // Set up routes
    router := api.SetupRouter()

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}