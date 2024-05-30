package middleware

import (
	"api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderRequest models.OrderRequest
		if err := c.ShouldBindJSON(&orderRequest); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		c.Next()
	}
}
