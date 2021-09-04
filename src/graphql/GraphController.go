package graphql

import (
	"github.com/friendsofgo/graphiql"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
)

func InitHandlers(myRouter *mux.Router) error {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	if err != nil {
		log.Printf("Initialize GraphQL Schema: %v\n", err)
		return nil
	}

	var graphHandler = handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	graphiHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		return err
	}

	myRouter.Handle("/graphql", graphHandler)
	myRouter.Handle("/graphiql", graphiHandler)

	return nil
}
