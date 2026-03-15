package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
	"github.com/pisondev/supply-management-api/internal/config"
	"github.com/pisondev/supply-management-api/internal/module/ingredient"
	"github.com/pisondev/supply-management-api/internal/module/inventory"
	"github.com/pisondev/supply-management-api/internal/module/supplier"
	"github.com/pisondev/supply-management-api/internal/module/warehouse"
	"github.com/pisondev/supply-management-api/utils"
)

func main() {
	log := utils.SetupLogger()

	err := godotenv.Load()
	if err != nil {
		log.Warn("failed to load .env file, falling back to system environment variables")
	}

	db := config.SetupDatabase(log)

	// Inisialisasi Modul Inventory
	inventoryRepo := inventory.NewRepository(db)
	inventoryService := inventory.NewService(inventoryRepo, db)
	inventoryController := inventory.NewController(inventoryService, log)

	// Inisialisasi Modul Ingredient
	ingredientRepo := ingredient.NewRepository(db)
	ingredientService := ingredient.NewService(ingredientRepo)
	ingredientController := ingredient.NewController(ingredientService, log)

	// Inisialisasi Modul Warehouse
	warehouseRepo := warehouse.NewRepository(db)
	warehouseService := warehouse.NewService(warehouseRepo)
	warehouseController := warehouse.NewController(warehouseService, log)

	// Inisialisasi Modul Supplier
	supplierRepo := supplier.NewRepository(db)
	supplierService := supplier.NewService(supplierRepo)
	supplierController := supplier.NewController(supplierService, log)

	// Fiber & Error Handler
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler(log),
	})

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Rate Limiter
	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(utils.WebResponse{
				Code:    fiber.StatusTooManyRequests,
				Status:  "error",
				Message: "too many request, please try again later",
			})
		},
	}))

	// Register Semua Routes
	api := app.Group("/api/v1")
	inventory.RegisterRoutes(api, inventoryController)
	ingredient.RegisterRoutes(api, ingredientController)
	warehouse.RegisterRoutes(api, warehouseController)
	supplier.RegisterRoutes(api, supplierController)

	app.Get("/health", func(c *fiber.Ctx) error {
		log.Info("health check endpoint accessed")
		return c.Status(fiber.StatusOK).JSON(utils.WebResponse{
			Code:    fiber.StatusOK,
			Status:  "success",
			Message: "supply management api is running properly",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Infof("starting api server on PORT %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
