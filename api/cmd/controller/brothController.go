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

// For this is a techincal test project, I wont support methods other than GET for broths and proteins
// I may expand it later on
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