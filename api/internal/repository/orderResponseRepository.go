package repository

import (
	"api/internal/database"
	"api/internal/models"
)

func GetAllOrderResponses() ([]models.OrderResponse, error) {
	var orderResponses []models.OrderResponse
	err := database.DB.Find(&orderResponses).Error
	return orderResponses, err
}

func GetOrderResponse(id string) (models.OrderResponse, error) {
	var orderResponse models.OrderResponse
	err := database.DB.Where("id=?", id).First(&orderResponse).Error
	return orderResponse, err
}

func CreateOrderResponse(o *models.OrderResponse) error {
	err := database.DB.Create(o).Error
	return err
}

func UpdateOrderResponse(o *models.OrderResponse) error {
	err := database.DB.Save(o).Error
	return err
}

func DeleteOrderResponse(o *models.OrderResponse) error {
	err := database.DB.Delete(o).Error
	return err
}