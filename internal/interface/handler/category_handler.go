package handler

import (
	"encoding/json"
	"net/http"
	"product-service/internal/app/model"
	"product-service/internal/app/service"
	"product-service/package/helper"
	"product-service/package/middleware"

	"github.com/go-playground/validator/v10"
)

type CategoryHandler struct {
	Service service.CategoryServiceInterface
}

func NewCategoryHandler(service service.CategoryServiceInterface) *CategoryHandler {
	return &CategoryHandler{Service: service}
}

// GetAllCategoriesHandler retrieves all categories
func (h *CategoryHandler) GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	categories, err := h.Service.GetAllCategories(ctx)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, "Failed to fetch categories", nil)
		return
	}

	middleware.WriteResponse(w, http.StatusOK, "Categories retrieved successfully", categories)
}

// CreateCategoryHandler handles the creation of a new category
func (h *CategoryHandler) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request model.SaveCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		middleware.WriteResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	// Validate the input
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		middleware.WriteResponse(w, http.StatusBadRequest, helper.GetMessageValidator(validate, err), nil)
		return
	}

	if err := h.Service.CreateCategory(ctx, request.Name); err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, "Failed to create category", nil)
		return
	}

	middleware.WriteResponse(w, http.StatusCreated, "Category created successfully", nil)
}
