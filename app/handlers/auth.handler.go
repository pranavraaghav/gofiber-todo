package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/pranavraagz/gofiber-todo/app/dal"
	"github.com/pranavraagz/gofiber-todo/app/types"
	"github.com/pranavraagz/gofiber-todo/utils"
	"github.com/pranavraagz/gofiber-todo/utils/jwt"
	"github.com/pranavraagz/gofiber-todo/utils/password"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	b := new(types.LoginDTO)

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	u := &types.UserResponse{}

	err := dal.FindUserByEmail(u, b.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email")
	}

	if err := password.Verify(u.Password, b.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid password")
	}

	t := jwt.Generate(&jwt.TokenPayload{ID: u.ID})

	return c.JSON(&types.AuthResponse{
		User: u,
		Auth: &types.AccessResponse{
			Token: t,
		},
	})
}

func Signup(c *fiber.Ctx) error {
	b := new(types.SignupDTO)

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	err := dal.FindUserByEmail(&struct{ ID string }{}, b.Email).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	user := &dal.User{
		Name:     b.Name,
		Password: password.Hash(b.Password),
		Email:    b.Email,
	}

	if err := dal.CreateUser(user); err.Error != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error.Error())
	}

	t := jwt.Generate(&jwt.TokenPayload{
		ID: user.ID,
	})

	return c.JSON(&types.AuthResponse{
		User: &types.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Auth: &types.AccessResponse{
			Token: t,
		},
	})
}
