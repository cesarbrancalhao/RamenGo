package controller

import (
	"api/internal/models"
	"api/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBroths(c *gin.Context) {

	res, err := repository.GetBroths()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateBroth(c *gin.Context) {
	var broth *models.Broth
	if err := c.ShouldBindJSON(&broth); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}	
	err := repository.CreateBroth(broth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, broth)
}