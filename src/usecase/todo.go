package usecase

import "hello-go-mock/src/entity"

type ToDoWriter interface {
	CreateToDo(t entity.ToDo) error
}
