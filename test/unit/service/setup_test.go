package service_test

import (
	"context"
	"flag"
	"os"
	"product-service/internal/app/repository"
	"product-service/internal/app/service"
	"product-service/internal/config"
	"product-service/internal/database"
	"product-service/package/logging"
	"testing"

	"github.com/google/uuid"
)

var (
	categoryRepo    repository.CategoryRepositoryInterface
	categoryService service.CategoryServiceInterface
	productRepo     repository.ProductRepositoryInterface
	productService  service.ProductServiceInterface
	ctx             context.Context
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {

	configFilePath := flag.String("config", "../../../cmd/config/config.yaml", "path to the config file")
	flag.Parse()

	// Load the configuration
	cfg, err := config.LoadConfig(*configFilePath)
	if err != nil {
		logging.Log.Fatalf("Error loading configuration: %v", err)
	}
	logging.Log.Infof("Load configuration from %v", *configFilePath)

	db, _ := database.InitDBPostgre(&cfg.Database.Main)

	redis, _ := database.InitRedis(&cfg.Redis)

	categoryRepo = repository.NewCategoryRepository(db)
	categoryService = service.NewCategoryService(categoryRepo)

	productRepo = repository.NewProductRepository(db, redis)
	productService = service.NewProductService(productRepo)

	requestID := uuid.New().String()
	ctx = context.WithValue(context.Background(), logging.RequestIDKey, requestID)
}
