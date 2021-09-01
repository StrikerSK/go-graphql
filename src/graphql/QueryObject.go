package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/strikersk/go-graphql/src/observer"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"todos": &graphql.Field{
			Type:        graphql.NewList(todoField),
			Description: "Read all todos",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				todos := observer.GetObserverInstance().FindAll()
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
		"done": &graphql.Field{
			Type:        todoField,
			Description: "Get todo by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
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
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"done": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)
