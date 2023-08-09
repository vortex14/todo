package main

import (
	"flag"

	"todo-list/internal/app"

	. "todo-list/internal/app/endpoints"
)

func main() {
	var portFlag = flag.Int("p", 3001, "port for app")

	flag.Parse()

	repository := (&app.InMemoryTodoRepository{}).Init()
	mainPoint := app.Endpoint{
		Service: &app.MainService{Repository: repository},
	}

	server := &app.Server{}
	server.Serve(&EndpointPing{})

	server.Serve(&EndpointCreateTask{
		Endpoint: mainPoint,
	})

	server.Serve(&EndpointUpdateTask{
		Endpoint: mainPoint,
	})

	server.Serve(&EndpointDeleteTask{
		Endpoint: mainPoint,
	})

	server.Serve(&EndpointGetTasks{
		Endpoint: mainPoint,
	})

	server.Start(*portFlag)
}
