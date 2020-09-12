package repository

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"hello-go-mock/src/entity"
	"hello-go-mock/src/usecase"
)

type ToDoWriterStdout struct{}

func NewToDoWriterStdout() usecase.ToDoWriter {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	return &ToDoWriterStdout{}
}

func (s *ToDoWriterStdout) CreateToDo(t entity.ToDo) error {
	data, err := json.Marshal(t)
	if err != nil {
		return errors.WithStack(err)
	}
	logrus.WithFields(logrus.Fields{
		"reciever": "ToDoWriterStdout",
		"method":   "CreateToDo",
	}).Info(string(data))
	return nil
}
