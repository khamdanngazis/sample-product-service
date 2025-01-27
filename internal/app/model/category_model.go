package model

type SaveCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}
