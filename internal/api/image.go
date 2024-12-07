package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ImageUpload handles image uploads
func ImageUpload(c *gin.Context) {
    file, _ := c.FormFile("file")
    log.Println("Uploaded File: ", file.Filename)
    
    // Save the file locally
    if err := c.SaveUploadedFile(file, "./uploads/"+file.Filename); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
