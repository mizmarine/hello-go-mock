//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock
package usecase

import "hello-go-mock/src/entity"

type ToDoWriter interface {
	CreateToDo(t entity.ToDo) error
}
