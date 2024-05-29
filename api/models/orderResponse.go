package models

type OrderResponse struct {
	Id string `json:"id"`
	Description string `json:"description"`
	Image string `json:"image"`
}