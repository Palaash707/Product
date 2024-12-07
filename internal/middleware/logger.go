package middleware

import (
	"time"

	"github.com/Palaash707/Product/internal/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LogRequest is a middleware that logs all requests and their response times
func LogRequest(c *gin.Context) {
    start := time.Now()

    // Process the request
    c.Next()

    // Log the request details after the response is sent
    duration := time.Since(start)

    logging.Logger.Info("Handled request",
        zap.String("method", c.Request.Method),
        zap.String("path", c.Request.URL.Path),
        zap.Int("status", c.Writer.Status()),
        zap.Float64("duration", duration.Seconds()),
    )
}
