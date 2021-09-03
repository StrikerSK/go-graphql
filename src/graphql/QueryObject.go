package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/strikersk/go-graphql/src/observer"
	"github.com/strikersk/go-graphql/src/types"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "query",
	Fields: graphql.Fields{
		"todos": &graphql.Field{
			Type:        graphql.NewList(todoField),
			Description: "Read all todos",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				todos, err := observer.GetObserverInstance().FindAll()
				if err != nil {
					return nil, err
				}

				return todos, nil
			},
		},
		"todo": &graphql.Field{
			Type:        todoField,
			Description: "Get todo by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				todoId, success := params.Args["id"].(string)
				todo, _ := observer.GetObserverInstance().FindByID(todoId)
				if success {
					return todo, nil
				}
				return todo, nil
			},
		},
		"getDone": &graphql.Field{
			Type:        graphql.NewList(todoField),
			Description: "Get every done todos",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				data, err := observer.GetObserverInstance().FindAll()
				if err != nil {
					return nil, err
				}

				var filteredTodos []types.Todo
				todos := data.([]types.Todo)
				for _, todo := range todos {
					if todo.Done {
						filteredTodos = append(filteredTodos, todo)
					}
				}

				return filteredTodos, nil
			},
		},
	},
})

//Type of field
//Need to define field
var todoField = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					todo := p.Source.(types.Todo)
					return todo.Name, nil
				},
			},
			"description": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					todo := p.Source.(types.Todo)
					return todo.Description, nil
				},
			},
			"subTasks": &graphql.Field{
				Type: graphql.NewList(graphql.NewObject(
					graphql.ObjectConfig{
						Name: "subTasks",
						Fields: graphql.Fields{
							"name": &graphql.Field{
								Type: graphql.String,
							},
							"description": &graphql.Field{
								Type: graphql.String,
							},
						},
					},
				)),
			},
			"done": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)
