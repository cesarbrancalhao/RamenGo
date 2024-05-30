package controller

import (
	"api/internal/models"
	"api/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProteins(c *gin.Context) {

	res, err := repository.GetProteins()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// For this is a techincal test project, I wont support methods other than GET for Proteins and proteins
// I may expand it later on
func CreateProtein(c *gin.Context) {
	var protein *models.Protein
	if err := c.ShouldBindJSON(&protein); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}	
	err := repository.CreateProtein(protein)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, protein)
}