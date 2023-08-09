package app

import (
	"fmt"
	"log"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	server *echo.Echo
	once   sync.Once
}

func (s *Server) up() {
	s.once.Do(func() {
		s.server = echo.New()

		s.server.Use(middleware.Logger())
		s.server.Use(middleware.Recover())
	})
}

func (s *Server) Serve(endpoint EndPointInterface) {
	s.up()
	s.server.Add(endpoint.GetMethod(), endpoint.GetPath(), endpoint.Callback)
}

func (s *Server) Start(port int) {
	log.Fatal(s.server.Start(fmt.Sprintf(":%d", port)))
}
