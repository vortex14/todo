package endpoints

import (
	"github.com/labstack/echo/v4"
	"todo-list/internal/app"
)

const UPDATE = "/task"

type EndpointUpdateTask struct {
	app.Endpoint
}

func (e *EndpointUpdateTask) GetMethod() string {
	return echo.PUT
}

func (e *EndpointUpdateTask) GetPath() string {
	return UPDATE
}

func (e *EndpointUpdateTask) Callback(c echo.Context) error {
	return e.Service.UpdateTask(c)
}
