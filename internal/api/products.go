package api

import "github.com/gin-gonic/gin"

func CreateProduct(c *gin.Context) {
    c.JSON(201, gin.H{"message": "Product created"})
}

func GetProduct(c *gin.Context) {
    c.JSON(200, gin.H{"product": "Product details"})
}

func ListProducts(c *gin.Context) {
    c.JSON(200, gin.H{"products": "List of products"})
}
