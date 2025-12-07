package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	_"github.com/99designs/gqlgen/handler"
	"github.com/99designs/gqlgen/graphql/handler"
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


	s, err := NewGraphQlServer(cfg.AccountServiceURL, cfg.CatalogServiceURL, cfg.OrderServiceURL)
	if err != nil {
		log.Fatalf("Failed to create GraphQL server: %v", err)
	}

	// s := server

	http.Handle("/graphql", handler.New(s.ToExecutableSchema()))
	http.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))
	log.Println("GraphQL server is running on http://localhost:8080/playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}