package repository

import (
	"api/internal/database"
	"api/internal/models"
)

func GetAllRecipes() ([]models.Recipe, error) {
	var Recipes []models.Recipe
	err := database.DB.Find(&Recipes).Error
	return Recipes, err
}

// For this is a techincal test project, I wont support the below methods
func GetRecipe(id string) (models.Recipe, error) {
	var recipe models.Recipe
	err := database.DB.Where("id=?", id).First(&recipe).Error
	return recipe, err
}

func CreateRecipe(o *models.Recipe) error {
	err := database.DB.Create(o).Error
	return err
}

func UpdateRecipe(o *models.Recipe) error {
	err := database.DB.Save(o).Error
	return err
}

func DeleteRecipe(o *models.Recipe) error {
	err := database.DB.Delete(o).Error
	return err
}