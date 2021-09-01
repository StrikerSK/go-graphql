package graphql

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/strikersk/go-graphql/src"
	"github.com/strikersk/go-graphql/src/observer"
	"log"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
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
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				todoId, _ := uuid.NewUUID()

				var newTodo src.Todo
				if err := mapstructure.Decode(params.Args, &newTodo); err != nil {
					log.Printf("GraphQL Create Todo: %v\n", err)
					return nil, err
				}

				_ = observer.GetObserverInstance().CreateData(newTodo)
				return todoId, nil
			},
		},
		"updateTodo": &graphql.Field{
			Type:        graphql.String,
			Description: "Create todo",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"done": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var newTodo src.Todo
				if err := mapstructure.Decode(params.Args, &newTodo); err != nil {
					log.Printf("GraphQL Update Todo: %v\n", err)
					return nil, err
				}

				if err := observer.GetObserverInstance().UpdateData(newTodo.Id, newTodo); err != nil {
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
				return true, nil
			},
		},
	},
})
