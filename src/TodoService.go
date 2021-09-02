package src

import (
	"errors"
	"github.com/google/uuid"
	"log"
)

const ServiceName = "TodoService"

type TodoService struct {
	todos []Todo
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
	todo, ok := input.(Todo)
	if !ok {
		return errors.New("input is not of type Todo")
	}

	if todo.Id == "" {
		todo.Id = uuid.NewString()
	}

	t.todos = append(t.todos, todo)
	return nil
}

func (t *TodoService) UpdateData(input interface{}) error {
	todo, ok := input.(Todo)
	if !ok {
		return errors.New("interface is not of type Todo")
	}

	todoIndex, err := t.findTodoIndex(todo.Id)

	if err != nil {
		log.Printf("Todo [%s] update: %v\n", todo.Id, err)
		return err
	}

	t.todos[todoIndex] = input.(Todo)
	return nil
}

func (t *TodoService) DeleteData(input interface{}) error {
	var filteredTodos []Todo

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
