package app

import "github.com/labstack/echo/v4"

type ToDoStorage interface {
	GetList() []*Task            // получение списка всех задач
	AddItem(task *Task) error    //создание новой задачи
	UpdateItem(task *Task) error //обновление существующей задачи
	DeleteItem(id string) error  // удаление задачи
	GetTask(id string) *Task
}

type EndPointInterface interface {
	GetPath() string
	GetMethod() string
	Callback(c echo.Context) error
}
