package main

import (
	"hello-go-mock/src/entity"
	"hello-go-mock/src/repository"
	"time"
)

func main() {
	repo := repository.NewToDoWriterStdout()
	todo := entity.ToDo{
		Title: "mock tamesu yo",
		Done:  false,
		Limit: time.Now(),
	}
	repo.CreateToDo(todo)
}
