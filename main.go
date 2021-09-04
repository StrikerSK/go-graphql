package main

import (
	"fmt"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/strikersk/go-graphql/src/graphql"
	"log"
	"os"

	"net/http"
)

func main() {
	myRouter := mux.NewRouter()

	if err := graphql.InitHandlers(myRouter); err != nil {
		log.Printf("Initialize GraphQL handlers: %v\n", err)
		os.Exit(1)
	}

	handler := cors.AllowAll().Handler(myRouter)
	fmt.Println(http.ListenAndServe(":5000", handler))
}
