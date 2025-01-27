package main

import (
	"flag"
	"os"
	"product-service/internal/app/repository"
	"product-service/internal/app/service"
	"product-service/internal/config"
	"product-service/internal/database"
	"product-service/internal/interface/handler"
	"product-service/internal/interface/router"
	"product-service/package/logging"
)

func main() {
	configFilePath := flag.String("config", "config/config.yaml", "path to the config file")
	//logFile := flag.String("log.file", "../logs", "Logging file")

	flag.Parse()

	initLogging()

	// Load the configuration
	cfg, err := config.LoadConfig(*configFilePath)
	if err != nil {
		logging.Log.Fatalf("Error loading configuration: %v", err)
	}

	logging.Log.Infof("Load configuration from %v", *configFilePath)

	db, err := database.InitDBPostgre(&cfg.Database.Main)

	if err != nil {
		logging.Log.Fatalf("Error initiate database connection: %v", err)
	}

	redis, err := database.InitRedis(&cfg.Redis)
	if err != nil {
		logging.Log.Fatalf("Error initiate database connection: %v", err)
	}

	if err != nil {
		logging.Log.Fatalf("Error initiate database connection: %v", err)
	}

	pingHandler := handler.NewPinghandler()

	//category
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	//product
	productRepo := repository.NewProductRepository(db, redis)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	httpRouter := router.NewMuxRouter()

	//ping handler
	httpRouter.GET("/api/v1/ping", pingHandler.Ping)

	//product handler
	httpRouter.POST("/api/v1/product", productHandler.AddProductHandler)
	httpRouter.GET("/api/v1/product", productHandler.GetProductListHandler)

	//category handler
	httpRouter.POST("/api/v1/category", categoryHandler.CreateCategoryHandler)
	httpRouter.GET("/api/v1/category", categoryHandler.GetAllCategoriesHandler)

	httpRouter.SERVE(cfg.AppPort)

}
func initLogging() {
	logging.InitLogger()
	logging.Log.SetOutput(os.Stdout)
}
