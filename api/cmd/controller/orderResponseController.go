package controller

import (
	"api/internal/models"
	"api/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrderResponses(c *gin.Context) {
	res, err := repository.GetAllOrderResponses()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}	
	c.JSON(http.StatusOK, res)
}


func GetOrderResponse(c *gin.Context) {
	id := c.Param("id")
	res, err := repository.GetOrderResponse(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateOrder(c *gin.Context) {
	id := c.Writer.Header().Get("recipe-id")
	rec, err := repository.GetRecipe(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	orderResponse := &models.OrderResponse{
		Id:          c.Writer.Header().Get("order-id"),
		Description: rec.Name,
		Image:       rec.Image,
	}

	err = repository.CreateOrderResponse(orderResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

    c.JSON(http.StatusOK, orderResponse)
}

func UpdateOrderResponse(c *gin.Context) {
	var orderResponse *models.OrderResponse
	if err := c.ShouldBindJSON(&orderResponse); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	err := repository.UpdateOrderResponse(orderResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderResponse)
}

func DeleteOrderResponse(c *gin.Context) {
	var orderResponse *models.OrderResponse
	if err := c.ShouldBindJSON(&orderResponse); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	err := repository.DeleteOrderResponse(orderResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderResponse)
}