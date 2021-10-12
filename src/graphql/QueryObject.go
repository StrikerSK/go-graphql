package graphql

import (
	"errors"
	"github.com/StrikerSK/go-graphql/src/observer"
	"github.com/StrikerSK/go-graphql/src/types"
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "query",
	Fields: graphql.Fields{
		"todos": &graphql.Field{
			Type: graphql.NewList(todoField),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				todos, err := observer.GetObserverInstance().FindAll()
				if err != nil {
					return nil, err
				}

				return todos, nil
			},
			Description: "Read all todos",
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
				if success {
					return nil, errors.New("cannot parse id value")
				}

				if todo, err := observer.GetObserverInstance().FindByID(todoId); err != nil {
					return nil, err
				} else {
					return todo, nil
				}
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
		Name: "todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					todo := p.Source.(types.Todo)
					return todo.Id, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					todo := p.Source.(types.Todo)
					return todo.Name, nil
				},
			},
			"description": &graphql.Field{
				Name: "description",
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
							"done": &graphql.Field{
								Type: graphql.Boolean,
							},
							"name": &graphql.Field{
								Type: graphql.String,
							},
							"description": &graphql.Field{
								Type: graphql.String,
							},
							"id": &graphql.Field{
								Type: graphql.String,
							},
						},
					},
				)),
			},
			"done": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					todo := p.Source.(types.Todo)
					return todo.Done, nil
				},
			},
		},
	},
)
