package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"hello-go-mock/src/entity"
	"hello-go-mock/src/usecase/mock"
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

func TestToDoService_SaveTodo_MultipleCall(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	writer := mock.NewMockToDoWriter(ctrl)

	svc := ToDoService{
		ToDoWriter: writer,
	}

	title1 := "test dayo1"
	title2 := "test dayo2"
	title3 := "test dayo3"

	c1 := writer.
		EXPECT().
		CreateToDo(gomock.Any()).
		Do(func(todo entity.ToDo) {
			assert.Equal(t, title1, todo.Title)
			assert.Equal(t, false, todo.Done)
			return
		}).
		Return(nil).
		Times(2)
	c2 := writer.
		EXPECT().
		CreateToDo(gomock.Any()).
		Do(func(todo entity.ToDo) {
			assert.Equal(t, title2, todo.Title)
			assert.Equal(t, false, todo.Done)
			return
		}).
		Return(nil).
		Times(3).
		After(c1)
	writer.
		EXPECT().
		CreateToDo(gomock.Any()).
		Do(func(todo entity.ToDo) {
			assert.Equal(t, title3, todo.Title)
			assert.Equal(t, false, todo.Done)
			return
		}).
		Return(nil).
		Times(1).
		After(c2)

	// call method multiple times
	svc.SaveTodo(title1)
	svc.SaveTodo(title1)
	svc.SaveTodo(title2)
	svc.SaveTodo(title2)
	svc.SaveTodo(title2)
	svc.SaveTodo(title3)
}
