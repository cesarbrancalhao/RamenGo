package models

import "github.com/go-playground/validator/v10"

type OrderResponse struct {
	Id          string `json:"id" validate:"nonzero"`
	Description string `json:"description" validate:"nonzero"`
	Image       string `json:"image" validate:"nonzero"`
}

var validateOrderResponse = validator.New()

func ValidateOrderResponse(o *OrderRequest) error {
	return validateOrderResponse.Struct(o)
}
