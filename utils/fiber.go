package utils

import "github.com/gofiber/fiber/v2"

func ParseBody(c *fiber.Ctx, body interface{}) *fiber.Error {
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ParseBody(ctx, body); err != nil {
		return err
	}
	return Validate(body)
}

func GetUser(c *fiber.Ctx) *uint {
	id, isFind := c.Locals("USER").(uint)
	if !isFind {
		return nil
	}
	return &id
}
