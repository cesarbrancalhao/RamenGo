package models

type OrderResponse struct {
	Id          int    `json:"id" validate:"gt=0" gorm:"primaryKey"`
	Description string `json:"description" validate:"nonzero"`
	Image       string `json:"image" validate:"nonzero"`
}
