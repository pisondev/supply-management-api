package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pisondev/supply-management-api/internal/config"
	"github.com/pisondev/supply-management-api/internal/module/inventory"
	"github.com/pisondev/supply-management-api/utils"
)

func main() {
	log := utils.SetupLogger()

	err := godotenv.Load()
	if err != nil {
		log.Warn("failed to load .env file, falling back to system environment variables")
	}

	db := config.SetupDatabase(log)

	inventoryRepo := inventory.NewRepository(db)
	inventoryService := inventory.NewService(inventoryRepo, db)
	inventoryController := inventory.NewController(inventoryService, log)

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler(log),
	})

	api := app.Group("/api/v1")
	inventory.RegisterRoutes(api, inventoryController)

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
