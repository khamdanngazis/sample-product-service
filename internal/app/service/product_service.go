package service

import (
	"context"
	"encoding/json"
	"fmt"
	"product-service/internal/app/entity"
	"product-service/internal/app/repository"
	"product-service/package/logging"
)

type ProductService struct {
	Repo repository.ProductRepositoryInterface
}

type ProductServiceInterface interface {
	AddProduct(ctx context.Context, product *entity.Product) error
	GetProductList(ctx context.Context, filter map[string]string, sort string) ([]entity.Product, error)
}

func NewProductService(repo repository.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductService{Repo: repo}
}

func (s *ProductService) AddProduct(ctx context.Context, product *entity.Product) error {
	// Save the new product to the database
	err := s.Repo.CreateProduct(product)
	if err != nil {
		logging.LogError(ctx, "Error CreateProduct: %v", err)
		return err
	}

	// Cache invalidation: Clear all relevant product cache
	err = s.invalidateProductCache(ctx)
	if err != nil {
		logging.LogError(ctx, "Error invalidateProductCache: %v", err)
	}

	return nil
}

func (s *ProductService) GetProductList(ctx context.Context, filter map[string]string, sort string) ([]entity.Product, error) {
	// Generate cache key based on filter and sort
	cacheKey := s.generateCacheKey(filter, sort)

	// Check cache first
	cachedProducts, err := s.Repo.GetCachedProducts(ctx, cacheKey)
	if err == nil && cachedProducts != nil {
		logging.LogInfo(ctx, "Cache hit for key: %v", cacheKey)
		return cachedProducts, nil
	} else if err != nil {
		logging.LogInfo(ctx, "Error GetCachedProducts: %v", err)
	}

	// Cache miss, fetch from DB
	products, err := s.Repo.GetProducts(filter, sort)
	if err != nil {
		logging.LogError(ctx, "Error GetProducts: %v", err)
		return nil, err
	}

	// Cache the result
	err = s.Repo.CacheProducts(ctx, cacheKey, products)
	if err != nil {
		logging.LogError(ctx, "Error CacheProducts: %v", err)
	}

	return products, nil
}

func (s *ProductService) generateCacheKey(filter map[string]string, sort string) string {
	key := "products_cache"
	if len(filter) > 0 {
		filterBytes, _ := json.Marshal(filter) // Convert filter to JSON string
		key += "_" + string(filterBytes)
	}
	if sort != "" {
		key += "_sort_" + sort
	}
	return key
}

func (s *ProductService) invalidateProductCache(ctx context.Context) error {
	cachePattern := "products_cache*"

	keys, err := s.Repo.GetCacheKeys(ctx, cachePattern)
	if err != nil {
		return err
	}

	fmt.Println("Keys to delete:", keys)

	// Delete all matching keys
	for _, key := range keys {
		if delErr := s.Repo.DeleteCacheKey(ctx, key); delErr != nil {
			logging.LogError(ctx, "Error DeleteCacheKey: %v", delErr)
		}
	}

	return nil
}
