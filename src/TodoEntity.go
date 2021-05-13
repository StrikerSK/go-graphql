package src

type Todo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ITodoService interface {
	readTodos() []Todo
}
