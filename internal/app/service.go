package app

import (
	"github.com/labstack/echo/v4"
)

type MainService struct {
	Repository *InMemoryTodoRepository
}

func (s *MainService) isValidNewTask(task *Task) bool {
	status := false

	if len(task.Title) > 0 && len(task.Description) > 0 {
		status = true
	}

	return status
}

func (s *MainService) isValidExistsTask(task *Task) bool {
	status := false

	if len(task.Title) > 0 &&
		len(task.Description) > 0 &&
		len(task.Id) > 0 {

		status = true
	}

	return status
}

func (s *MainService) CreateTask(c echo.Context) error {
	_task := &Task{}
	err := c.Bind(_task)

	if err != nil {
		return c.JSON(400, &Response{Message: NotValidRequest.Error()})
	}

	if !s.isValidNewTask(_task) {
		return c.JSON(400, &Response{Message: RequiredCreateFieldErr.Error()})
	}

	err = s.Repository.AddItem(_task)

	if err != nil {
		return c.JSON(400, &Response{Message: err.Error()})
	}

	return c.JSON(200, _task)

}

func (s *MainService) UpdateTask(c echo.Context) error {
	_task := &Task{}
	err := c.Bind(_task)

	if err != nil {
		return c.JSON(400, &Response{Message: NotValidRequest.Error()})
	}

	if !s.isValidExistsTask(_task) {
		return c.JSON(400, &Response{Message: RequiredUpdateFieldErr.Error()})
	}

	err = s.Repository.UpdateItem(_task)

	if err != nil {
		return c.JSON(400, &Response{Message: err.Error()})
	}

	return c.JSON(200, _task)

}

func (s *MainService) DeleteTask(c echo.Context) error {
	_task := &Task{}
	err := c.Bind(_task)

	if err != nil {
		return c.JSON(400, &Response{Message: NotValidRequest.Error()})
	}

	if len(_task.Id) == 0 {
		return c.JSON(400, &Response{Message: RequiredIdFieldErr.Error()})
	}

	err = s.Repository.DeleteItem(_task.Id)

	if err != nil {
		return c.JSON(400, &Response{Message: err.Error()})
	}

	return c.JSON(200, &Response{Message: "Ok"})

}

func (s *MainService) GetTasks(c echo.Context) error {

	return c.JSON(200, s.Repository.GetList())

}
