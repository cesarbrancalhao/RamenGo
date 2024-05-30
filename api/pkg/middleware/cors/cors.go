package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
    w := c.Writer
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-api-key")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET")

    apiKeyHeader := http.CanonicalHeaderKey("x-api-key")
    if _, ok := c.Request.Header[apiKeyHeader];!ok || len(c.Request.Header[apiKeyHeader]) == 0 {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing API key"})
        return
    }
	
	if c.Request.Header[apiKeyHeader][0] != os.Getenv("API_KEY") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

    c.Next()
}