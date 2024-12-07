package main

import (
	"log"

	"github.com/Palaash707/Product/internal/api"
	"github.com/Palaash707/Product/internal/db"
)

func main() {
    db.ConnectDB()

    router := api.SetupRouter()
    log.Println("Server running at http://localhost:8080")
    router.Run(":8080")
}
