package repository

import (
	"encoding/json"
	"fmt"
	"hello-go-mock/src/entity"
)

type toDoWriterStdout struct{}

func NewToDoWriterStdout() *toDoWriterStdout {
	return &toDoWriterStdout{}
}

func (s *toDoWriterStdout) CreateToDo(t entity.ToDo) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
