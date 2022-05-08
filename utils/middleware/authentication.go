package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pranavraagz/gofiber-todo/utils/jwt"
)

func Auth(c *fiber.Ctx) error {
	h := c.Get("Authorization")

	if h == "" {
		return fiber.ErrUnauthorized
	}

	chunks := strings.Split(h, " ")

	if len(chunks) > 2 {
		return fiber.ErrUnauthorized
	}

	user, err := jwt.Verify(chunks[1])

	if err != nil {
		return fiber.ErrUnauthorized
	}

	c.Locals("USER", user.ID)

	return c.Next()
}
