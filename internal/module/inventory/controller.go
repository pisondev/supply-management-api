package inventory

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pisondev/supply-management-api/utils"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	service Service
	log     *logrus.Logger
}

func NewController(service Service, log *logrus.Logger) *Controller {
	return &Controller{service, log}
}

func (ctrl *Controller) RecordMovement(c *fiber.Ctx) error {
	var request RecordMovementRequest
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	ctrl.log.Info("processing inventory movement request")
	result, err := ctrl.service.RecordMovement(&request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(utils.WebResponse{
		Code:    fiber.StatusCreated,
		Status:  "success",
		Message: "inventory movement recorded successfully",
		Data:    result,
	})
}
