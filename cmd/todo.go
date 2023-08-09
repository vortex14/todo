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
	mainService := &app.MainService{Repository: repository}

	server := &app.Server{}
	server.Serve(&EndpointPing{})

	server.Serve(&EndpointCreateTask{
		Endpoint: app.Endpoint{
			Service: mainService,
		},
	})

	server.Serve(&EndpointUpdateTask{
		Endpoint: app.Endpoint{
			Service: mainService,
		},
	})

	server.Serve(&EndpointDeleteTask{
		Endpoint: app.Endpoint{
			Service: mainService,
		},
	})

	server.Serve(&EndpointGetTasks{
		Endpoint: app.Endpoint{
			Service: mainService,
		},
	})

	server.Start(*portFlag)
}
