package graphql

import (
	"github.com/StrikerSK/go-graphql/src/observer"
	"github.com/StrikerSK/go-graphql/src/types"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"log"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutation",
	Fields: graphql.Fields{
		"createTodo": &graphql.Field{
			Type:        graphql.String,
			Description: "Create todo",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"subTasks": &graphql.ArgumentConfig{
					Type: graphql.NewList(subTaskObject),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var newTodo types.Todo
				if err := mapstructure.Decode(params.Args, &newTodo); err != nil {
					log.Printf("GraphQL Create Todo: %v\n", err)
					return nil, err
				}

				newTodo.Id = uuid.NewString()
				if err := observer.GetObserverInstance().CreateData(newTodo); err != nil {
					return nil, err
				}

				return newTodo.Id, nil
			},
		},
		"updateTodo": &graphql.Field{
			Type:        graphql.String,
			Description: "Update todo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"subTasks": &graphql.ArgumentConfig{
					Type: graphql.NewList(subTaskObject),
				},
				"done": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var newTodo types.Todo
				if err := mapstructure.Decode(params.Args, &newTodo); err != nil {
					log.Printf("GraphQL Update Todo: %v\n", err)
					return nil, err
				}

				if err := observer.GetObserverInstance().UpdateData(newTodo); err != nil {
					return nil, err
				}

				return "Todo updated", nil
			},
		},
		"deleteTodo": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Delete todo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				todoID := params.Args["id"].(string)

				if err := observer.GetObserverInstance().DeleteData(todoID); err != nil {
					return false, err
				}

				return true, nil
			},
		},
	},
})

//GraphQL's object for inputting nested structure
var subTaskObject = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Sub Tasks",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
