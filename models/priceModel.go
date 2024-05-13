package models

type Price struct {
	Size     string  `json:"size" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Currency string  `json:"currency" validate:"required"`
}
