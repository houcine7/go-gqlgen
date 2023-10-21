package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/houcine7/graphql-server/graph"
	database "github.com/houcine7/graphql-server/internal/db"
	auth "github.com/houcine7/graphql-server/internal/middleware"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	
	
	
	database.InitDb()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", auth.AuthMiddleware()(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
