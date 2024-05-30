package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Protein struct {
	gorm.Model
	ImageInactive string `json:"imageInactive" validate:"url"`
	ImageActive string `json:"imageActive" validate:"url"`
	Name string `json:"name" validate:"nonzero"`
	Description string `json:"description" validate:"nonzero"`
	Price int `json:"price" validate:"gt=0"`
}

var validateProtein = validator.New()

func ValidateProtein(p *Protein) error {
	return validateProtein.Struct(p)
}