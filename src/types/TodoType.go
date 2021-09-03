package types

type Todo struct {
	Id       string `json:"id"`
	Task     `mapstructure:",squash"`
	SubTasks []Task `json:"subTasks"`
	Done     bool   `json:"done"`
}

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
