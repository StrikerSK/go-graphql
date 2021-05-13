package src

var (
	customTodo = Todo{
		Name:        "Test Todo",
		Description: "Test Todo",
		Done:        false,
	}
)

var todos []Todo

func FindAll() []Todo {
	return []Todo{customTodo, customTodo, customTodo}
}

func FindById(todoID uint) Todo {
	return customTodo
}
