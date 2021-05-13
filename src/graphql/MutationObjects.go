package graphql

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/strikersk/go-graphql/src"
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
				todoName, _ := params.Args["name"].(string)
				todoDescription, _ := params.Args["description"].(string)
				todoId, _ := uuid.NewUUID()

				createTodo := src.Todo{
					Id:          todoId.String(),
					Name:        todoName,
					Description: todoDescription,
				}

				src.CreateTodo(createTodo)
				return todoId, nil
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
				todoId, _ := params.Args["id"].(string)
				_, present := src.FindById(todoId)
				return present, nil
			},
		},
	},
})
