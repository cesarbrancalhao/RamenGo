package middleware

import (
	"api/internal/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func ValidateGeneratedId(c *gin.Context, orderId string) {

	brothIdRegex := regexp.MustCompile(`^[0-9]+$`)

	if !brothIdRegex.MatchString(orderId) {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid orderId"})
	}
}