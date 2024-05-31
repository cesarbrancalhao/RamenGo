package middleware

import (
	"api/internal/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func ValidateOrderRequest(c *gin.Context, order *models.OrderRequest) {

	brothIdRegex := regexp.MustCompile(`^[1-3]+$`)
	proteinIdRegex := regexp.MustCompile(`^[1-3]+$`)

	if !brothIdRegex.MatchString(order.BrothId) {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid brothId"})
	}
	if !proteinIdRegex.MatchString(order.ProteinId) {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid proteinId"})
	}
}