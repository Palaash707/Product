package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Palaash707/Product/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProductsPagination(t *testing.T) {
    router := gin.Default()
    router.GET("/products", ListProducts)

    // Create a product for testing
    db.CreateProduct(1, "Product 1", "Description", []string{"image1.jpg"}, 100.0)

    // Send a request with pagination parameters
    req, _ := http.NewRequest("GET", "/products?page=1&limit=1", nil)
    resp := performRequest(router, req)

    // Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }

    // Convert the body to a string
    bodyString := string(body)

    assert.Equal(t, http.StatusOK, resp.StatusCode)
    assert.Contains(t, bodyString, "Product 1")
}

func performRequest(router http.Handler, req *http.Request) *http.Response {
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    return w.Result()
}