package service

import (
	"errors"
	"github.com/StrikerSK/go-graphql/src/types"
	"github.com/google/uuid"
	"log"
)

const ServiceName = "TodoService"

type TodoService struct {
	todos []types.Todo
}

func (t *TodoService) GetName() string {
	return ServiceName
}

func (t *TodoService) FindAll() (interface{}, error) {
	return t.todos, nil
}

func (t *TodoService) FindByID(todoID interface{}) (interface{}, error) {
	todoIndex, err := t.findTodoIndex(todoID.(string))
	if err != nil {
		log.Printf("Todo [%s] read: %v\n", todoID, err)
		return nil, err
	}

	return t.todos[todoIndex], nil
}

func (t *TodoService) CreateData(input interface{}) error {
	todo, ok := input.(types.Todo)
	if !ok {
		return errors.New("input is not of type Todo")
	}

	if todo.Id == "" {
		todo.Id = uuid.NewString()
	}

	for index := range todo.SubTasks {
		todo.SubTasks[index].GenerateID()
	}

	t.todos = append(t.todos, todo)
	return nil
}

func (t *TodoService) UpdateData(input interface{}) error {
	todo, ok := input.(types.Todo)
	if !ok {
		return errors.New("interface is not of type Todo")
	}

	todoIndex, err := t.findTodoIndex(todo.Id)

	if err != nil {
		log.Printf("Todo [%s] update: %v\n", todo.Id, err)
		return err
	}

	t.todos[todoIndex] = input.(types.Todo)
	return nil
}

func (t *TodoService) DeleteData(input interface{}) error {
	var filteredTodos []types.Todo

	for _, item := range t.todos {
		if item.Id != input {
			filteredTodos = append(filteredTodos, item)
		}
	}

	t.todos = filteredTodos
	return nil
}

func (t *TodoService) findTodoIndex(todoID string) (int, error) {
	for index, todo := range t.todos {
		if todo.Id == todoID {
			return index, nil
		}
	}

	return 0, errors.New("data not found")
}
