package controller

import (
	"api/internal/models"
	"api/internal/repository"
	"errors"
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

// For this is a techincal test project, I wont support the below methods, I built them for showcase purposes
func CreateBroth(c *gin.Context) {
	var broth *models.Broth
	if err := c.ShouldBindJSON(&broth); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}	
	// err := repository.CreateBroth(broth)
	err := errors.New("service unavailable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, broth)
}

func UpdateBroth(c *gin.Context) {
	var broth *models.Broth
	if err := c.ShouldBindJSON(&broth); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}	
	// err := repository.UpdateBroth(broth)
	err := errors.New("service unavailable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, broth)
}

func DeleteBroth(c *gin.Context) {
	var broth *models.Broth
	if err := c.ShouldBindJSON(&broth); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}	
	// err := repository.DeleteBroth(broth)
	err := errors.New("service unavailable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, broth)
}