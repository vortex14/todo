package endpoints

import (
	"github.com/labstack/echo/v4"
	"todo-list/internal/app"
)

const (
	LIST = "/list"
)

type EndpointGetTasks struct {
	app.Endpoint
}

func (e *EndpointGetTasks) GetMethod() string {
	return echo.GET
}

func (e *EndpointGetTasks) GetPath() string {
	return LIST
}

func (e *EndpointGetTasks) Callback(c echo.Context) error {
	return e.Service.GetTasks(c)
}
