package warehouse

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pisondev/supply-management-api/utils"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	service Service
	log     *logrus.Logger
}

func NewController(s Service, l *logrus.Logger) *Controller { return &Controller{s, l} }

func (ctrl *Controller) Create(c *fiber.Ctx) error {
	var req CreateWarehouseRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}
	res, err := ctrl.service.Create(&req)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(utils.WebResponse{Code: 201, Status: "success", Data: res})
}

func (ctrl *Controller) GetAll(c *fiber.Ctx) error {
	res, err := ctrl.service.GetAll()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(utils.WebResponse{Code: 200, Status: "success", Data: res})
}

func (ctrl *Controller) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	res, err := ctrl.service.GetByID(id)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(utils.WebResponse{Code: 200, Status: "success", Data: res})
}

func (ctrl *Controller) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req UpdateWarehouseRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}
	res, err := ctrl.service.Update(id, &req)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(utils.WebResponse{Code: 200, Status: "success", Data: res})
}

func (ctrl *Controller) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := ctrl.service.Delete(id); err != nil {
		return err
	}
	return c.Status(200).JSON(utils.WebResponse{Code: 200, Status: "success"})
}
