package service_test

import (
	"product-service/internal/app/entity"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProduct(t *testing.T) {
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

	productRepo.CreateProduct(&product)
	//create filter and sortetr
	filter := map[string]string{
		"id": strconv.Itoa(int(product.ID)),
	}
	//get product
	products, err := productService.GetProductList(ctx, filter, "id")
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	assert.Equal(t, product.Name, products[0].Name)
	assert.Equal(t, product.Price, products[0].Price)
	assert.Equal(t, product.CategoryID, products[0].CategoryID)
	//delete product test

	productRepo.DeleteProduct(product.ID)
	categoryRepo.DeleteCategory(category.ID)
}
