package utils

import "github.com/gofiber/fiber/v2"

func ParseBody(c *fiber.Ctx, body interface{}) *fiber.Error {
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

func ParseBodyAndValidate(c *fiber.Ctx, body interface{}) *fiber.Error {
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

func GetUser(c *fiber.Ctx) *uint {
	id, _ := c.Locals("USER").(uint)
	return &id
}
