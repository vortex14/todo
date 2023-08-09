package endpoints

import (
	"github.com/labstack/echo/v4"
	"todo-list/internal/app"
)

const PING = "/"

type EndpointPing struct {
	app.Endpoint
}

func (e *EndpointPing) GetMethod() string {
	return echo.GET
}

func (e *EndpointPing) GetPath() string {
	return PING
}

func (e *EndpointPing) Callback(c echo.Context) error {
	return c.JSON(200, &app.Response{Message: "Ok"})
}
