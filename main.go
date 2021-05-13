package main

import (
	"context"
	"fmt"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	graphHandler "github.com/graphql-go/handler"
	"github.com/rs/cors"
	"github.com/strikersk/go-graphql/src/graphql"

	"net/http"
)

func httpHeaderMiddleware(next *graphHandler.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "token", r.Header.Get("Token"))

		next.ContextHandler(ctx, w, r)
	})
}

func main() {
	myRouter := mux.NewRouter()

	myRouter.Handle("/graphql", httpHeaderMiddleware(graphql.GraphHandler))
	myRouter.Handle("/graphiql", graphql.GraphiQLHandler)

	handler := cors.AllowAll().Handler(myRouter)
	fmt.Println(http.ListenAndServe(":5000", handler))
}
