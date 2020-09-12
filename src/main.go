package main

import "hello-go-mock/src/entity"

func main() {

}

type ToDoWriter interface {
	CreateToDo(t entity.ToDo) error
}

