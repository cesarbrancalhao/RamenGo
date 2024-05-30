package models

type Broth struct {
	Id            int    `json:"id" validate:"gt=0" gorm:"primaryKey"`
	ImageInactive string `json:"imageInactive" validate:"url"`
	ImageActive   string `json:"imageActive" validate:"url"`
	Name          string `json:"name" validate:"nonzero"`
	Description   string `json:"description" validate:"nonzero"`
	Price         int    `json:"price" validate:"gt=0"`
}