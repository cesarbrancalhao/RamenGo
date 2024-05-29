package models

import "github.com/go-playground/validator/v10"

type Broth struct {
	Id            string `json:"id" validate:"nonzero"`
	ImageInactive string `json:"imageInactive" validate:"url"`
	ImageActive   string `json:"imageActive" validate:"url"`
	Name          string `json:"name" validate:"nonzero"`
	Description   string `json:"description" validate:"nonzero"`
	Price         int    `json:"price" validate:"gt=0"`
}

var validateBroth = validator.New()

func ValidateBroth(b *Broth) error {
	return validateBroth.Struct(b)
}
