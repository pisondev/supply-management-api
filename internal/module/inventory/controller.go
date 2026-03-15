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

func (ctrl *Controller) GetStocks(c *fiber.Ctx) error {
	var filter StockFilterParam
	if err := c.QueryParser(&filter); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid query parameters")
	}

	ctrl.log.Info("fetching inventory stocks")
	result, err := ctrl.service.GetStocks(&filter)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(utils.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "stocks retrieved successfully",
		Data:    result,
	})
}

func (ctrl *Controller) GetMovements(c *fiber.Ctx) error {
	var filter MovementFilterParam
	if err := c.QueryParser(&filter); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid query parameters")
	}

	ctrl.log.Info("fetching inventory movements")
	result, err := ctrl.service.GetMovements(&filter)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(utils.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "movements retrieved successfully",
		Data:    result,
	})
}
