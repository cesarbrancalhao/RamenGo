package models

type OrderResponse struct {
	Id          string `json:"id" validate:"nonzero" gorm:"primaryKey"`
	Description string `json:"description" validate:"nonzero"`
	Image       string `json:"image" validate:"nonzero"`
}
