package endpoints

import (
	"github.com/labstack/echo/v4"
	"todo-list/internal/app"
)

const CREATE = "/task"

type EndpointCreateTask struct {
	app.Endpoint
}

func (e *EndpointCreateTask) GetMethod() string {
	return echo.POST
}

func (e *EndpointCreateTask) GetPath() string {
	return CREATE
}

func (e *EndpointCreateTask) Callback(c echo.Context) error {
	return e.Service.CreateTask(c)
}
