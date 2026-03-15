package warehouse

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router, ctrl *Controller) {
	group := router.Group("/warehouses")
	group.Post("/", ctrl.Create)
	group.Get("/", ctrl.GetAll)
	group.Get("/:id", ctrl.GetByID)
	group.Put("/:id", ctrl.Update)
	group.Delete("/:id", ctrl.Delete)
}
