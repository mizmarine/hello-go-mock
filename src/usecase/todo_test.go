package usecase

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"hello-go-mock/src/entity"
	"hello-go-mock/src/usecase/mock"
	"testing"
)

func TestToDoService_SaveTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	writer := mock.NewMockToDoWriter(ctrl)

	svc := ToDoService{
		ToDoWriter: writer,
	}

	writer.
		EXPECT().
		CreateToDo(gomock.Any()).
		Do(func(t entity.ToDo) {
			fmt.Println(t)
			return
		}).
		Return(nil)

	svc.SaveTodo("test dayo")
}
