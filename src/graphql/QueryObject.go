package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/strikersk/go-graphql/src"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"todos": &graphql.Field{
			Type:        graphql.NewList(todoField),
			Description: "Read all todos",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return src.FindAll(), nil
			},
		},
		"todo": &graphql.Field{
			Type:        todoField,
			Description: "Get todo by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, success := params.Args["id"].(int)
				if success {
					return src.FindById(uint(id)), nil
				}
				return nil, nil
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
				id, success := params.Args["id"].(int)
				if success {
					return src.FindById(uint(id)), nil
				}
				return nil, nil
			},
		},
	},
})

//Type of field
//Need to define field
var todoField = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Exercise",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
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
