package controller

import (
	"api/internal/models"
	"api/internal/repository"
	"errors"
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

// For this is a techincal test project, I wont support the below methods, I still built them for showcase purposes
func CreateProtein(c *gin.Context) {
	var protein *models.Protein
	if err := c.ShouldBindJSON(&protein); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}	
	// err := repository.CreateProtein(protein)
	err := errors.New("service unavailable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, protein)
}

func UpdateProtein(c *gin.Context) {
	var protein *models.Protein
	if err := c.ShouldBindJSON(&protein); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}	
	// err := repository.UpdateProtein(protein)
	err := errors.New("service unavailable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, protein)
}

func DeleteProtein(c *gin.Context) {
	var protein *models.Protein
	if err := c.ShouldBindJSON(&protein); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}	
	// err := repository.DeleteProtein(protein)
	err := errors.New("service unavailable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, protein)
}