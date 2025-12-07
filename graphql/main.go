package main

import "log"

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


	NewGraphQlServer(cfg.AccountServiceURL, cfg.CatalogServiceURL, cfg.OrderServiceURL)
	if err != nil {
		log.Fatalf("Failed to create GraphQL server: %v", err)
	}


	http.Handler("graphql", )
}