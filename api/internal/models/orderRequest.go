package models

import "github.com/go-playground/validator/v10"

type OrderRequest struct {
	BrothId   string `json:"brothId" validate:"nonzero"`
	ProteinId string `json:"proteinId" validate:"nonzero"`
}

var validateOrderRequest = validator.New()

func ValidateOrderRequest(o *OrderRequest) error {
	return validateOrderRequest.Struct(o)
}
