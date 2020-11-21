package usecase

import (
	"github.com/golang/mock/gomock"
	"hello-go-mock/src/entity"
	"hello-go-mock/src/usecase/mock"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestToDoService_SaveTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	writer := mock.NewMockToDoWriter(ctrl)

	svc := ToDoService{
		ToDoWriter: writer,
	}

	title := "test dayo"

	writer.
		EXPECT().
		CreateToDo(gomock.Any()).
		Do(func(todo entity.ToDo) {
			assert.Equal(t, title, todo.Title)
			assert.Equal(t, false, todo.Done)
			return
		}).
		Return(nil).
		Times(1)

	svc.SaveTodo(title)
}
