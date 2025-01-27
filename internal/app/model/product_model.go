package model

type SaveProductRequest struct {
	Name       string  `json:"name" validate:"required"`
	CategoryID uint    `json:"category_id" validate:"required"`
	Price      float64 `json:"price" validate:"required"`
}
