package app

import "errors"

var (
	TaskNotFound           = errors.New("task not found")
	NotValidRequest        = errors.New("not valid requests")
	RequiredCreateFieldErr = errors.New("title, description are required")
	RequiredUpdateFieldErr = errors.New("title, description, id are required")
	RequiredIdFieldErr     = errors.New("title, description, id are required")
	UndefinedErr           = errors.New("undefined error")
)
