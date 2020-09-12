package main

import (
	"hello-go-mock/src/repository"
	"hello-go-mock/src/usecase"
)

func main() {
	svc := usecase.ToDoService{
		ToDoWriter: repository.NewToDoWriterStdout(),
	}
	svc.SaveTodo("task dayo")
}
