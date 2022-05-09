package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pranavraagz/gofiber-todo/app/dal"
	"github.com/pranavraagz/gofiber-todo/app/types"
	"github.com/pranavraagz/gofiber-todo/utils"
)

func CreateTodo(c *fiber.Ctx) error {
	b := new(types.CreateDTO)

	if err := utils.ParseBody(c, b); err != nil {
		return err
	}

	d := &dal.Todo{
		User: utils.GetUser(c),
		Task: b.Task,
	}

	if err := dal.CreateTodo(d).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.TodoCreateResponse{
		Todo: &types.TodoResponse{
			ID:        d.ID,
			Task:      d.Task,
			Completed: d.Completed,
		},
	})
}

func CheckTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	b := new(types.CheckTodoDTO)
	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	if err := dal.UpdateTodo(todoID, utils.GetUser(c), map[string]interface{}{"completed": b.Completed}).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Msg: "Todo updated successfully",
	})
}

func GetTodos(c *fiber.Ctx) error {
	d := &[]types.TodoResponse{}

	err := dal.FindTodosByUser(d, utils.GetUser(c)).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.TodosResponse{
		Todos: d,
	})

}

func UpdateTodo(c *fiber.Ctx) error {
	b := new(types.CreateDTO)
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}
	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	err := dal.UpdateTodo(todoID, utils.GetUser(c), &dal.Todo{
		Task: b.Task,
	}).Error

	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Msg: "Updated successfully",
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	err := dal.DeleteTodo(todoID, utils.GetUser(c)).Error
	if err != nil {
		return err
	}

	return c.JSON(&types.MsgResponse{
		Msg: "Deleted successfully",
	})
}
