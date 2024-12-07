package api

import (
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    router.POST("/products", CreateProduct)
    router.GET("/products/:id", GetProduct)
    router.GET("/products", ListProducts)

    return router
}
