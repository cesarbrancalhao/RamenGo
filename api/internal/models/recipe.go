package models

type Recipe struct {
	Id      int    `json:"id" validate:"gt=0" gorm:"primaryKey"`
	Name 	string `json:"name" validate:"nonzero"`
	Image	string `json:"image" validate:"nonzero"`
}