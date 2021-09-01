package src

import (
	"log"
)

type LoggingTodoService struct {
}

func (t *LoggingTodoService) GetName() string {
	return "LoggingTodoService"
}

func (t *LoggingTodoService) FindAll() interface{} {
	log.Println("User requested finding all Todos")
	return nil
}

func (t *LoggingTodoService) FindByID(todoID interface{}) (interface{}, error) {
	log.Println("User requested finding Todo by ID: ", todoID)
	return nil, nil
}

func (t *LoggingTodoService) CreateData(input interface{}) error {
	log.Println("User requested creating Todo: ", input)
	return nil
}

func (t *LoggingTodoService) UpdateData(id interface{}, input interface{}) error {
	log.Println("User requested updating Todo with ID: ", id, " with Data: ", input)
	return nil
}

func (t *LoggingTodoService) DeleteData(input interface{}) error {
	log.Println("User requested creating Todo: ", input)
	return nil
}
