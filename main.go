package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	utils "github.com/pisondev/supply-management-api/pkg"
)

func main() {
	log := utils.SetupLogger()

	err := godotenv.Load()
	if err != nil {
		log.Warn("failed to load .env file, back to system environment variables")
	}

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		log.Info("health check endpoint accessed")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "supply management api is running properly",
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
