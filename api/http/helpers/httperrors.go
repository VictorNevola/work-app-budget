package helpers

import "github.com/gofiber/fiber/v2"

type ApiError struct {
	Status  int
	Message string
}

func HandlerError(c *fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(&ApiError{
		Status:  statusCode,
		Message: err.Error(),
	})
}
