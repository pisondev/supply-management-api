package supplier

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

func (ctrl *Controller) Create(c *fiber.Ctx) error {
	var request CreateSupplierRequest
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	ctrl.log.Info("creating new supplier")
	result, err := ctrl.service.Create(&request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(utils.WebResponse{
		Code:    fiber.StatusCreated,
		Status:  "success",
		Message: "supplier created successfully",
		Data:    result,
	})
}

func (ctrl *Controller) GetAll(c *fiber.Ctx) error {
	ctrl.log.Info("fetching all suppliers")
	result, err := ctrl.service.GetAll()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(utils.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "suppliers retrieved successfully",
		Data:    result,
	})
}

func (ctrl *Controller) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	ctrl.log.Info("fetching supplier by id: ", id)

	result, err := ctrl.service.GetByID(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(utils.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "supplier retrieved successfully",
		Data:    result,
	})
}

func (ctrl *Controller) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var request UpdateSupplierRequest
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	ctrl.log.Info("updating supplier: ", id)
	result, err := ctrl.service.Update(id, &request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(utils.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "supplier updated successfully",
		Data:    result,
	})
}

func (ctrl *Controller) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	ctrl.log.Info("deleting supplier: ", id)
	err := ctrl.service.Delete(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(utils.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "supplier deleted successfully",
		Data:    nil,
	})
}
