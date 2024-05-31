package repository

import (
	"api/internal/database"
	"api/internal/models"
)

func GetBroths() ([]models.Broth, error) {
	var broths []models.Broth
	err := database.DB.Find(&broths).Error
	return broths, err
}

// For this is a techincal test project, I wont support the below methods
func CreateBroth(b *models.Broth) error {
	err := database.DB.Create(b).Error
	return err
}

func UpdateBroth(b *models.Broth) error {
	err := database.DB.Save(b).Error
	return err
}

func DeleteBroth(b *models.Broth) error {
	err := database.DB.Delete(b).Error
	return err
}