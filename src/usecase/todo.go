//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock
package usecase

import (
	"time"

	"hello-go-mock/src/entity"
)

type ToDoWriter interface {
	CreateToDo(t entity.ToDo) error
}

type ToDoService struct {
	ToDoWriter ToDoWriter
}

func (s *ToDoService) SaveTodo(title string) error {
	todo := entity.ToDo{
		Title: title,
		Done:  false,
		Limit: time.Now(),
	}
	return s.ToDoWriter.CreateToDo(todo)
}
