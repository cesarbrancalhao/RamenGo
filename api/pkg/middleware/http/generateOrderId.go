package middleware

import (
	"api/internal/models"
	"encoding/json"
	"io"
	"net/http"
	"os"

	validator "api/pkg/middleware/validations"

	"github.com/gin-gonic/gin"
)

type OrderIDResponse struct {
    OrderId string `json:"orderId"`
}

func GenerateId(c *gin.Context) {
	client := &http.Client{}
    req, err := http.NewRequest("POST", "https://api.tech.redventures.com.br/orders/generate-id", nil)
    if err!= nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("x-api-key", os.Getenv("API_KEY"))

    resp, err := client.Do(req)
    if err!= nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
    }
    defer resp.Body.Close()

    responseBody, err := io.ReadAll(resp.Body)
    if err!= nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
    }

    var response OrderIDResponse
	if err := json.Unmarshal(responseBody, &response); err!= nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
	}	

	validator.ValidateGeneratedId(c, response.OrderId)
    c.Writer.Header().Set("order-id", response.OrderId)

    c.Next()
}