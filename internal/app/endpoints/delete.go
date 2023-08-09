package endpoints

import (
	"github.com/labstack/echo/v4"
	"todo-list/internal/app"
)

const DELETE = "/task"

type EndpointDeleteTask struct {
	app.Endpoint
}

func (e *EndpointDeleteTask) GetMethod() string {
	return echo.DELETE
}

func (e *EndpointDeleteTask) GetPath() string {
	return DELETE
}

func (e *EndpointDeleteTask) Callback(c echo.Context) error {
	return e.Service.DeleteTask(c)
}
