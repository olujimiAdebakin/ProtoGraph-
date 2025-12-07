// package main

// import (
// 	"log"
// 	"net/http"

	
// 	"github.com/99designs/gqlgen/graphql/playground"
// 	_"github.com/99designs/gqlgen/handler"
// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/kelseyhightower/envconfig"
// )

// type AppConfig struct {
// 	AccountServiceURL string `envconfig:"ACCOUNT_SERVICE_URL" default:"http://localhost:8081"`
// 	CatalogServiceURL string `envconfig:"CATALOG_SERVICE_URL" default:"http://localhost:8082"`
// 	OrderServiceURL   string `envconfig:"ORDER_SERVICE_URL" default:"http://localhost:8083"`
// }

// func main() {
// 	var cfg AppConfig
// 	err := envconfig.Process("", &cfg)
// 	if err != nil {
// 		log.Fatalf("Failed to process envconfig: %v", err)
// 	}


// 	s, err := NewGraphQlServer(cfg.AccountServiceURL, cfg.CatalogServiceURL, cfg.OrderServiceURL)
// 	if err != nil {
// 		log.Fatalf("Failed to create GraphQL server: %v", err)
// 	}

// 	// s := server

// 	http.Handle("/graphql", handler.New(s.ToExecutableSchema()))
// 	http.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))
// 	log.Println("GraphQL server is running on http://localhost:8080/playground")
	
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

package main

import (
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountServiceURL string `envconfig:"ACCOUNT_SERVICE_URL" default:"http://localhost:8081"`
	CatalogServiceURL string `envconfig:"CATALOG_SERVICE_URL" default:"http://localhost:8082"`
	OrderServiceURL   string `envconfig:"ORDER_SERVICE_URL" default:"http://localhost:8083"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Failed to process envconfig: %v", err)
	}

	// Initialize your GraphQL server, capturing both server instance and error
	s, err := NewGraphQlServer(cfg.AccountServiceURL, cfg.CatalogServiceURL, cfg.OrderServiceURL)
	if err != nil {
		log.Fatalf("Failed to create GraphQL server: %v", err)
	}

	// Create executable schema from your server resolvers
	execSchema := s.ToExecutableSchema()

	// Setup the GraphQL handler with the executable schema
	srv := handler.New(execSchema)

	// Register the GraphQL endpoint
	http.Handle("/graphql", srv)

	// Register Playground UI at /playground for easy testing
	http.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))

	log.Println("GraphQL server is running on http://localhost:8080/playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
