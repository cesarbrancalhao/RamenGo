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

// For this is a techincal test project, I wont support the below methods, I built them for showcase purposes
func CreateRecipe(c *gin.Context) {
	var Recipe *models.Recipe
	if err := c.ShouldBindJSON(&Recipe); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	// err := repository.CreateRecipe(Recipe)
	err := errors.New("service unavailable, to make an Recipe use /recipe")
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
	// err := repository.UpdateRecipe(Recipe)
	err := errors.New("service unavailable")
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
	// err := repository.DeleteRecipe(Recipe)
	err := errors.New("service unavailable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Recipe)
}