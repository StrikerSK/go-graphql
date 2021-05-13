package src

import "log"

var (
	customTodo = Todo{
		Id:          "Non existing todo",
		Name:        "Non existing todo",
		Description: "Non existing todo",
		Done:        false,
	}
)

var todos []Todo

func FindAll() []Todo {
	return todos
}

func FindById(todoID string) (Todo, bool) {
	todoIndex, present := findTodoIndex(todoID)

	if !present {
		log.Printf("Cannot find todo with id %s\n", todoID)
		return customTodo, present
	}

	return todos[todoIndex], present
}

func CreateTodo(createdTodo Todo) {
	todos = append(todos, createdTodo)
	return
}

func UpdateTodo(updatedTodo Todo) {
	todoIndex, present := findTodoIndex(updatedTodo.Id)

	if !present {
		log.Printf("Cannot update todo with id %s\n", updatedTodo.Id)
		return
	}

	todos[todoIndex] = updatedTodo
	return
}

func findTodoIndex(todoID string) (int, bool) {
	for index, todo := range todos {
		if todo.Id == todoID {
			return index, true
		}
	}

	return 0, false
}
