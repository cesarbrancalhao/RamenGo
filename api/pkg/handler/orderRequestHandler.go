package handler

import (
	"api/common/filter"
	"api/internal/models"
	"net/http"

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
	
	orderRecipeId := filter.FindOrderRecipeId(brothId, proteinId)
	c.Writer.Header().Set("recipe-id", orderRecipeId)

	c.Next()
}