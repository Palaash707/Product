package db

import (
	"log"

	"github.com/lib/pq" // Import pq to handle PostgreSQL array conversion
)

// Product represents a product in the system
type Product struct {
    ID                    int      `json:"id"`
    ProductName           string   `json:"product_name"`
    ProductDescription    string   `json:"product_description"`
    ProductImages         []string `json:"product_images"`
    CompressedProductImages []string `json:"compressed_product_images"`
    ProductPrice          float64  `json:"product_price"`
}

func CreateProduct(userID int, productName, productDescription string, productImages []string, productPrice float64) (int, error) {
    // SQL query to insert the product into the products table
    query := `
        INSERT INTO products (user_id, product_name, product_description, product_images, product_price)
        VALUES ($1, $2, $3, $4, $5) RETURNING id
    `
    
    var id int
    // Using pq.Array() to convert the Go slice of strings into a PostgreSQL array
    err := DB.QueryRow(query, userID, productName, productDescription, pq.Array(productImages), productPrice).Scan(&id)
    if err != nil {
        log.Printf("Error inserting product: %v", err)
        return 0, err
    }

    return id, nil
}

// GetProductByID retrieves a product by its ID from the database
func GetProductByID(id int) (*Product, error) {
    query := `
        SELECT id, product_name, product_description, product_images, compressed_product_images, product_price
        FROM products WHERE id = $1
    `
    
    var product Product
    err := DB.QueryRow(query, id).Scan(
        &product.ID, &product.ProductName, &product.ProductDescription,
        pq.Array(&product.ProductImages), pq.Array(&product.CompressedProductImages), &product.ProductPrice,
    )
    
    if err != nil {
        log.Printf("Error retrieving product: %v", err)
        return nil, err
    }

    return &product, nil
}

// ListProducts retrieves all products from the database
func ListProducts(offset, limit int) ([]Product, error) {
    query := `
        SELECT id, product_name, product_description, product_images, compressed_product_images, product_price
        FROM products
        LIMIT $1 OFFSET $2
    `
    
    rows, err := DB.Query(query, limit, offset)
    if err != nil {
        log.Printf("Error fetching products: %v", err)
        return nil, err
    }
    defer rows.Close()

    var products []Product
    for rows.Next() {
        var product Product
        err := rows.Scan(
            &product.ID, &product.ProductName, &product.ProductDescription,
            pq.Array(&product.ProductImages), pq.Array(&product.CompressedProductImages), &product.ProductPrice,
        )
        if err != nil {
            log.Printf("Error reading product: %v", err)
            return nil, err
        }
        products = append(products, product)
    }

    return products, nil
}