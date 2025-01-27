package service_test

import (
	"product-service/internal/app/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {
	//create category test
	category := entity.Category{
		Name: "Category test",
	}
	categoryRepo.CreateCategory(&category)
	//create product test
	product := entity.Product{
		Name:       "Product test",
		CategoryID: category.ID,
		Price:      1000,
	}

	err := productRepo.CreateProduct(&product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)

	productRepo.DeleteProduct(product.ID)
	categoryRepo.DeleteCategory(category.ID)
}
