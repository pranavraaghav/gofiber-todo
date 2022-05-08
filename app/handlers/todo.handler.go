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
