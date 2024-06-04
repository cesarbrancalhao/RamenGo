package middleware

import (
	"api/internal/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Cors(c *gin.Context) {
    w := c.Writer
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-api-key")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

    if c.Request.Method == http.MethodOptions {
        c.AbortWithStatus(http.StatusOK)
        return
    }

    apiKeyHeader := http.CanonicalHeaderKey("x-api-key")
    if _, ok := c.Request.Header[apiKeyHeader];!ok || len(c.Request.Header[apiKeyHeader]) == 0 {
        c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Missing API key"})
        return
    }

	if bcrypt.CompareHashAndPassword([]byte(c.Request.Header[apiKeyHeader][0]), []byte(os.Getenv("API_KEY"))) != nil {
        c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid API key"})
        return
	}

    c.Next()
}