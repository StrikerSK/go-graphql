package graphql

import (
	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
	//Mutation: rootMutation,
})

var GraphHandler = handler.New(&handler.Config{
	Schema: &schema,
	Pretty: true,
})

var GraphiQLHandler, _ = graphiql.NewGraphiqlHandler("/graphql")
