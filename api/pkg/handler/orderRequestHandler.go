package handler

import (
	"api/internal/models"
	"net/http"
	"strconv"

	validator "api/pkg/middleware/validations"

	"github.com/gin-gonic/gin"
)
func HandleOrderRequest(c *gin.Context) {
	var order *models.OrderRequest
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	validator.ValidateOrderRequest(c, order)

	proteinId := order.ProteinId
	brothId := order.BrothId
	
	orderResponseId := FindOrderResponseId(brothId, proteinId)
	c.Writer.Header().Set("order-id", orderResponseId)

	c.Next()
}

func FindOrderResponseId(b string, p string) string {
	broth := map[string]int { "1": 0, "2": 3, "3": 6 }
	protein := map[string]int { "1": 1, "2": 2, "3": 3 }
	
	orderId := strconv.Itoa(broth[b] + protein[p])
	return orderId
}