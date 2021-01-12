package todos

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/pallat/todos/logger"
)

type Task struct {
	gorm.Model
	Task      string
	Processed bool
}

func (Task) TableName() string {
	return "todos"
}

func NewNewTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		db.AutoMigrate(Task{})

		var todo struct {
			Task string `json:"task"`
		}

		logger := logger.Extract(c)
		logger.Info("new task todo........")

		if err := c.Bind(&todo); err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": errors.Wrap(err, "new task").Error(),
			})
		}

		if err := db.Create(&Task{
			Task: todo.Task,
		}).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "create task").Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{})
	}
}



func GetTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		db.AutoMigrate(Task{})

		logger := logger.Extract(c)
		logger.Info("get task todo........")

		var todo []Task

		if err := db.Find(&todo).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "get task").Error(),
			})
		}

		return c.JSON(http.StatusOK, todo)
	}
}


func UpdateTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		db.AutoMigrate(Task{})

		logger := logger.Extract(c)
		logger.Info("new task todo........ %")

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "update task get id").Error(),
			})
		}

		var todo Task

		if err := db.Model(&todo).Where("id = ?",id).Update("processed",true).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "update task").Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{})
	}
}


func DeleteTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		db.AutoMigrate(Task{})

		logger := logger.Extract(c)
		logger.Info("new task todo........ %")

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "delete task get id").Error(),
			})
		}

		var todo Task

		if err := db.Where("id = ?",id).Delete(&todo).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "update task").Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{})
	}
}


