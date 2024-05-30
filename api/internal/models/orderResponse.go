package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OrderResponse struct {
	gorm.Model
	Description string `json:"description" validate:"nonzero"`
	Image       string `json:"image" validate:"nonzero"`
}

var validateOrderResponse = validator.New()

func ValidateOrderResponse(o *OrderResponse) error {
	return validateOrderResponse.Struct(o)
}
