package repository

import (
	"context"
	"encoding/json"
	"product-service/internal/app/entity"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

type ProductRepositoryInterface interface {
	CreateProduct(product *entity.Product) error
	GetProducts(filter map[string]string, sort string) ([]entity.Product, error)
	DeleteProduct(id uint) error
	CacheProducts(ctx context.Context, key string, products []entity.Product) error
	GetCachedProducts(ctx context.Context, key string) ([]entity.Product, error)
	GetCacheKeys(ctx context.Context, pattern string) ([]string, error)
	DeleteCacheKey(ctx context.Context, key string) error
}

func NewProductRepository(db *gorm.DB, redisClient *redis.Client) ProductRepositoryInterface {
	return &ProductRepository{DB: db, Redis: redisClient}
}

func (r *ProductRepository) CreateProduct(product *entity.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) GetProducts(filter map[string]string, sort string) ([]entity.Product, error) {
	var products []entity.Product
	db := r.DB.Preload(clause.Associations)

	// Filter by product name
	if name, ok := filter["name"]; ok && name != "" {
		db = db.Where("products.name ILIKE ?", "%"+name+"%")
	}

	// Filter by product ID
	if id, ok := filter["id"]; ok && id != "" {
		db = db.Where("products.id = ?", id)
	}

	if category, ok := filter["category"]; ok && category != "" {
		db = db.Joins("LEFT JOIN categories AS Category ON products.category_id = Category.id").
			Where("Category.name ILIKE ?", "%"+category+"%")
	}

	order := "ASC" // Default sorting order
	if len(sort) > 4 && sort[len(sort)-4:] == "_asc" {
		sort = sort[:len(sort)-4]
		order = "ASC"
	} else if len(sort) > 5 && sort[len(sort)-5:] == "_desc" {
		sort = sort[:len(sort)-5]
		order = "DESC"
	}

	switch sort {
	case "price":
		db = db.Order("products.price " + order)
	case "name":
		db = db.Order("products.name " + order)
	case "date":
		db = db.Order("products.created_at " + order)
	default:
		db = db.Order("products.id " + order)
	}

	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	return r.DB.Delete(&entity.Product{}, id).Error
}

func (r *ProductRepository) CacheProducts(ctx context.Context, key string, products []entity.Product) error {
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}
	return r.Redis.Set(ctx, key, data, time.Hour).Err()
}

func (r *ProductRepository) GetCachedProducts(ctx context.Context, key string) ([]entity.Product, error) {
	var products []entity.Product
	data, err := r.Redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			// Cache miss
			return nil, nil
		}
		// Log Redis error
		//log.Printf("Redis error: %v", err)
		return nil, err
	}

	// Unmarshal JSON data
	if err := json.Unmarshal([]byte(data), &products); err != nil {
		//log.Printf("Failed to unmarshal products from cache: %v", err)
		return nil, err
	}

	// Handle empty data (optional)
	if len(products) == 0 {
		return nil, nil
	}

	return products, nil
}

func (r *ProductRepository) GetCacheKeys(ctx context.Context, pattern string) ([]string, error) {
	// Use Redis SCAN command to fetch keys matching the pattern
	iter := r.Redis.Scan(ctx, 0, pattern, 0).Iterator()
	var keys []string
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return keys, nil
}

func (r *ProductRepository) DeleteCacheKey(ctx context.Context, key string) error {
	return r.Redis.Del(ctx, key).Err()
}
