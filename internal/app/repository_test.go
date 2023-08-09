package app

import (
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInMemoryTodoRepository(t *testing.T) {

	repo := (&InMemoryTodoRepository{}).Init()
	_uuid, _ := uuid.NewUUID()

	Convey("add a new task", t, func(c C) {

		newTask := &Task{Id: _uuid.String(), Description: "some description", Title: "some title"}

		err := repo.AddItem(newTask)

		c.So(err, ShouldBeNil)

		_task := repo.GetTask(_uuid.String())

		c.So(_task.Id, ShouldEqual, _uuid.String())

	})

	Convey("remove task", t, func(c C) {

		errD := repo.DeleteItem(_uuid.String())

		c.So(errD, ShouldBeNil)

		errD = repo.DeleteItem(_uuid.String())

		c.So(errD, ShouldEqual, TaskNotFound)

	})

	Convey("update task", t, func(c C) {

		newTask := &Task{Id: _uuid.String(), Description: "some description", Title: "some title"}

		err := repo.AddItem(newTask)
		c.So(err, ShouldBeNil)

		err = repo.UpdateItem(&Task{Id: _uuid.String(), Title: "new title"})

		c.So(err, ShouldBeNil)

		_task := repo.GetTask(_uuid.String())

		c.So(_task.Title, ShouldEqual, "new title")

		err = repo.UpdateItem(&Task{Id: "not found id", Title: "new title"})

		c.So(err, ShouldEqual, TaskNotFound)

	})

}
