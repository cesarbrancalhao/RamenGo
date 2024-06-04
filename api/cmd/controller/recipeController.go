package controller

import (
	"api/internal/models"
	"api/internal/repository"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllRecipes(c *gin.Context) {
	res, err := repository.GetAllRecipes()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}	
	c.JSON(http.StatusOK, res)
}

func GetRecipe(c *gin.Context) {
	id := c.Param("id")
	res, err := repository.GetRecipe(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// For this is a technical test project, I wont support the below methods on the deployed API, I built them for showcase purposes
func CreateRecipe(c *gin.Context) {
	var Recipe *models.Recipe
	if err := c.ShouldBindJSON(&Recipe); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	err := repository.CreateRecipe(Recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Recipe)
}

func UpdateRecipe(c *gin.Context) {
	var Recipe *models.Recipe
	if err := c.ShouldBindJSON(&Recipe); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	err := repository.UpdateRecipe(Recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Recipe)
}

func DeleteRecipe(c *gin.Context) {
	var Recipe *models.Recipe
	if err := c.ShouldBindJSON(&Recipe); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	err := repository.DeleteRecipe(Recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Recipe)
}