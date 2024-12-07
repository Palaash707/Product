package api

import (
	"log"
	"net/http"

	"strconv"

	"github.com/Palaash707/Product/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// CreateProduct handles creating a new product
func CreateProduct(c *gin.Context) {
    var product struct {
        UserID             int      `json:"user_id"`
        ProductName        string   `json:"product_name"`
        ProductDescription string   `json:"product_description"`
        ProductImages      []string `json:"product_images"`
        ProductPrice       float64  `json:"product_price"`
    }

    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Log the incoming data for debugging
    log.Printf("Received product data: %+v\n", product)

    // Check if user exists
    var userExists bool
    userCheckQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`
    err := db.DB.QueryRow(userCheckQuery, product.UserID).Scan(&userExists)
    
    if err != nil {
        log.Printf("Error checking if user exists: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user existence", "details": err.Error()})
        return
    }

    if !userExists {
        log.Printf("User with ID %d does not exist", product.UserID)
        c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
        return
    }

    // Call the db function to insert the product
    id, err := db.CreateProduct(product.UserID, product.ProductName, product.ProductDescription, product.ProductImages, product.ProductPrice)
    if err != nil {
        log.Printf("Error creating product: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product", "details": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"id": id})
}
//get product
func GetProduct(c *gin.Context) {
    id := c.Param("id")

    query := `
        SELECT id, product_name, product_description, product_images, compressed_product_images, product_price
        FROM products WHERE id = $1
    `

    // Define the product struct
    var product struct {
        ID                 int      `json:"id"`
        ProductName        string   `json:"product_name"`
        ProductDescription string   `json:"product_description"`
        ProductImages      []string `json:"product_images"`
        CompressedImages   []string `json:"compressed_product_images"`
        ProductPrice       float64  `json:"product_price"`
    }

    // Use pq.Array() to scan PostgreSQL array fields into Go slices
    err := db.DB.QueryRow(query, id).Scan(
        &product.ID, &product.ProductName, &product.ProductDescription,
        pq.Array(&product.ProductImages), pq.Array(&product.CompressedImages), &product.ProductPrice,
    )

    if err != nil {
        log.Printf("Error retrieving product: %v", err)
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    // Return the product in JSON format
    c.JSON(http.StatusOK, product)
}

// ListProducts retrieves all products from the database
func ListProducts(c *gin.Context) {
    // Get query parameters for pagination
    pageStr := c.DefaultQuery("page", "1")   // Default to page 1 if not provided
    limitStr := c.DefaultQuery("limit", "10") // Default to 10 products per page

    // Convert strings to integers
    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
        return
    }
    limit, err := strconv.Atoi(limitStr)
    if err != nil || limit < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
        return
    }

    // Calculate OFFSET for pagination
    offset := (page - 1) * limit

    // Fetch the products from the database with pagination
    products, err := db.ListProducts(offset, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
        return
    }

    c.JSON(http.StatusOK, products)
}