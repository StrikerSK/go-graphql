package src

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

func FindById(todoID string) Todo {
	for _, todo := range todos {
		if todo.Id == todoID {
			return todo
		}
	}

	return customTodo
}

func CreateTodo(createdTodo Todo) {
	todos = append(todos, createdTodo)
	return
}
