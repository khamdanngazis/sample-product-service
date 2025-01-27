package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"product-service/internal/app/entity"
	"product-service/internal/config"
	"product-service/internal/database"
	"product-service/package/logging"
	"time"
)

func main() {
	configFilePath := flag.String("config", "config/config.yaml", "path to the config file")
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
		logging.Log.Fatalf("Error initiating database connection: %v", err)
	}

	err = db.AutoMigrate(
		&entity.Category{},
		&entity.Product{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully!")

	// Create categories
	categories := []entity.Category{
		{Name: "Sayuran"},
		{Name: "Protein"},
		{Name: "Buah"},
		{Name: "Snack"},
		{Name: "Minuman"},
	}

	// Seed categories and update their IDs
	for i := range categories {
		if err := db.Create(&categories[i]).Error; err != nil {
			log.Fatalf("Failed to create category: %v", err)
		}
	}
	log.Println("Categories seeded successfully!")

	// Reload categories to ensure IDs are set correctly
	if err := db.Find(&categories).Error; err != nil {
		log.Fatalf("Failed to retrieve categories: %v", err)
	}

	// Sample data for product names
	productNames := map[string][]string{
		"Sayuran": {"Bayam", "Brokoli", "Wortel", "Selada", "Kangkung", "Sawi"},
		"Protein": {"Daging Ayam", "Daging Sapi", "Ikan Salmon", "Udang", "Telur"},
		"Buah":    {"Apel", "Jeruk", "Pisang", "Anggur", "Mangga", "Semangka"},
		"Snack":   {"Keripik Kentang", "Cokelat", "Biskuit", "Kacang", "Permen"},
		"Minuman": {"Teh Botol", "Jus Jeruk", "Air Mineral", "Kopi", "Soda"},
	}

	// Create 100 products
	for i := 1; i <= 100; i++ {
		category := categories[rand.Intn(len(categories))] // Random category
		name := productNames[category.Name][rand.Intn(len(productNames[category.Name]))] + fmt.Sprintf(" %d", i)
		price := float64(rand.Intn(20000-500) + 500) // Random price between 500 and 20,000

		product := entity.Product{
			Name:       name,
			CategoryID: category.ID, // Use the correct category ID
			Price:      price,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		if err := db.Create(&product).Error; err != nil {
			log.Fatalf("Failed to create product: %v", err)
		}
	}

	log.Println("Products seeded successfully!")
}

func initLogging() {
	logging.InitLogger()
	logging.Log.SetOutput(os.Stdout)
}
