package main

import (
	"fmt"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/strikersk/go-graphql/src/graphql"

	"net/http"
)

func main() {
	myRouter := mux.NewRouter()

	myRouter.Handle("/graphql", graphql.GraphHandler)
	myRouter.Handle("/graphiql", graphql.GraphiQLHandler)

	handler := cors.AllowAll().Handler(myRouter)
	fmt.Println(http.ListenAndServe(":5000", handler))
}
