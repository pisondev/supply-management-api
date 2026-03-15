package ingredient

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router, ctrl *Controller) {
	group := router.Group("/ingredients")

	group.Post("/", ctrl.Create)
	group.Get("/", ctrl.GetAll)
	group.Get("/:id", ctrl.GetByID)
	group.Put("/:id", ctrl.Update)
	group.Delete("/:id", ctrl.Delete)
}
