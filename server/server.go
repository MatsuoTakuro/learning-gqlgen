package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MatsuoTakuro/learning-gqlgen/graph/generated"
	"github.com/MatsuoTakuro/learning-gqlgen/graph/resolver"
)

const defaultPort = "8085"
const title = "gqlgen-todo"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(resolver.New()))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		// send this panic somewhere
		log.Print(err)
		debug.PrintStack()
		return errors.New("user message on panic")
	})

	http.Handle("/", playground.Handler(title, "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for %s", port, title)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
