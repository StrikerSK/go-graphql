package src

import (
	"errors"
	"github.com/google/uuid"
	"log"
)

const ServiceName = "TodoService"

var (
	customTodo = Todo{
		Id:          "Non existing todo",
		Name:        "Non existing todo",
		Description: "Non existing todo",
		Done:        false,
	}
)

type TodoService struct {
	todos []Todo
}

func (t *TodoService) GetName() string {
	return ServiceName
}

func (t *TodoService) FindAll() interface{} {
	return t.todos
}

func (t *TodoService) FindByID(todoID interface{}) (interface{}, error) {
	todoIndex, present := t.findTodoIndex(todoID.(string))

	if !present {
		log.Printf("Cannot find todo with id %s\n", todoID)
		return customTodo, nil
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

func (t *TodoService) UpdateData(id interface{}, input interface{}) error {
	todoIndex, present := t.findTodoIndex(id.(string))

	if !present {
		log.Printf("Todo [%s] update: value not found\n", id)
		return errors.New("value not found")
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

func (t *TodoService) findTodoIndex(todoID string) (int, bool) {
	for index, todo := range t.todos {
		if todo.Id == todoID {
			return index, true
		}
	}

	return 0, false
}
