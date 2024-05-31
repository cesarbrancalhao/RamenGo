package controller

import (
	"api/internal/models"
	"api/internal/repository"
	"errors"
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

func GetOrderResponseByRequest(c *gin.Context) {
	id := c.Writer.Header().Get("order-id")
	res, err := repository.GetOrderResponse(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// For this is a techincal test project, I wont support the below methods, I built them for showcase purposes
func CreateOrderResponse(c *gin.Context) {
	var orderResponse *models.OrderResponse
	if err := c.ShouldBindJSON(&orderResponse); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	// err := repository.CreateOrderResponse(orderResponse)
	err := errors.New("service unavailable, to make an order use /order")
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
	// err := repository.UpdateOrderResponse(orderResponse)
	err := errors.New("service unavailable")
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
	// err := repository.DeleteOrderResponse(orderResponse)
	err := errors.New("service unavailable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderResponse)
}