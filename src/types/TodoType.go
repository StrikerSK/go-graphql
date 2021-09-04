package types

import "github.com/google/uuid"

type Todo struct {
	Task     `mapstructure:",squash"`
	SubTasks []Task `json:"subTasks"`
}

type Task struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (t *Task) GenerateID() {
	t.Id = uuid.NewString()
}
