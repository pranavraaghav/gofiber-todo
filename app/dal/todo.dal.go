package dal

import (
	"github.com/pranavraagz/gofiber-todo/config/database"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Task      string `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	User      *uint  `gorm:"not null" gorm:"index"`
}

func CreateTodo(todo *Todo) *gorm.DB {
	return database.DB.Create(todo)
}

func FindTodo(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&Todo{}).Take(dest, conds)
}

func FindTodoByUser(dest interface{}, todoIden interface{}, userIden interface{}) *gorm.DB {
	return FindTodo(dest, "id = ? AND user - ?", todoIden, userIden)
}

func FindTodosByUser(dest interface{}, userIden interface{}) *gorm.DB {
	return database.DB.Model(&Todo{}).Find(dest, "user = ?", userIden)
}

func DeleteTodo(todoIden interface{}, userIden interface{}) *gorm.DB {
	return database.DB.Unscoped().Delete(&Todo{}, "id = ? AND user = ?", todoIden, userIden)
}

func UpdateTodo(todoIden interface{}, userIden interface{}, data interface{}) *gorm.DB {
	return database.DB.Model(&Todo{}).Where("id = ? AND user = ?", todoIden, userIden).Updates(data)
}
