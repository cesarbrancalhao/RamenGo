package repository

import (
	"api/internal/database"
	"api/internal/models"
	"errors"
)

func GetProteins() ([]models.Protein, error) {
	var proteins []models.Protein
	err := database.DB.Find(&proteins).Error
	return proteins, err
}

// For this is a techincal test project, I wont support the below methods
// I may expand it later on
func CreateProtein(b *models.Protein) error {
	// err := database.DB.Create(b).Error
	// return err
	return errors.New("service unavailable")
}

func UpdateProtein(b *models.Protein) error {
	err := database.DB.Save(b).Error
	return err
}

func DeleteProtein(b *models.Protein) error {
	err := database.DB.Delete(b).Error
	return err
}