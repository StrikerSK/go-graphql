package service

import (
	"github.com/StrikerSK/go-graphql/src/types"
	"log"
)

type LoggingTodoService struct{}

func (t *LoggingTodoService) GetName() string {
	return "LoggingTodoService"
}

func (t *LoggingTodoService) FindAll() (interface{}, error) {
	log.Println("User requested finding all Todos")
	return nil, nil
}

func (t *LoggingTodoService) FindByID(todoID interface{}) (interface{}, error) {
	log.Println("User requested finding Todo by ID: ", todoID)
	return nil, nil
}

func (t *LoggingTodoService) CreateData(input interface{}) error {
	log.Println("User requested creating Todo: ", input)
	return nil
}

func (t *LoggingTodoService) UpdateData(input interface{}) error {
	todo := input.(types.Todo)
	log.Println("User requested updating Todo with ID: ", todo.Id, " with Data: ", todo)
	return nil
}

func (t *LoggingTodoService) DeleteData(input interface{}) error {
	log.Println("User requested creating Todo: ", input)
	return nil
}
