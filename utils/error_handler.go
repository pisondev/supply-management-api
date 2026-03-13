package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(log *logrus.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		status := "internal_server_error"
		message := err.Error()

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			status = "client_error"
			message = e.Message
		}

		log.Errorf("api error: %v", err)

		return ctx.Status(code).JSON(WebResponse{
			Code:    code,
			Status:  status,
			Message: message,
		})
	}
}
