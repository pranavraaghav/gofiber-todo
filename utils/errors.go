package utils

import "github.com/gofiber/fiber/v2"

type httpError struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	e, ok := err.(*fiber.Error)
	if ok {
		code = e.Code
	}

	return c.Status(code).JSON(&httpError{
		StatusCode: code,
		Error:      err.Error(),
	})
}
