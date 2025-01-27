package handler

import (
	"encoding/json"
	"net/http"
	"product-service/internal/app/entity"
	"product-service/internal/app/model"
	"product-service/internal/app/service"
	"product-service/package/helper"
	"product-service/package/middleware"

	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	Service service.ProductServiceInterface
}

func NewProductHandler(service service.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) GetProductListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := map[string]string{}
	for key := range r.URL.Query() {
		filter[key] = r.URL.Query().Get(key)
	}
	sort := r.URL.Query().Get("sort")

	products, err := h.Service.GetProductList(ctx, filter, sort)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, "Failed to fetch products", nil)
		return
	}

	middleware.WriteResponse(w, http.StatusOK, "Products retrieved successfully", products)
}

func (h *ProductHandler) AddProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request model.SaveProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		middleware.WriteResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		middleware.WriteResponse(w, http.StatusBadRequest, helper.GetMessageValidator(validate, err), nil)
		return
	}

	product := entity.Product{
		Name:       request.Name,
		Price:      request.Price,
		CategoryID: request.CategoryID,
	}

	if err := h.Service.AddProduct(ctx, &product); err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, "Failed to add product", nil)
		return
	}

	middleware.WriteResponse(w, http.StatusCreated, "Product added successfully", nil)
}
