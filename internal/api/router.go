package api

import (
	"github.com/Palaash707/Product/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Apply JWT authentication middleware
    router.Use(middleware.AuthMiddleware())

    router.POST("/products", CreateProduct)
    router.GET("/products/:id", GetProduct)
    router.GET("/products", ListProducts)
    return router
}