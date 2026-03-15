package inventory

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router, ctrl *Controller) {
	inventoryGroup := router.Group("/inventory")
	inventoryGroup.Post("/movements", ctrl.RecordMovement)
}
