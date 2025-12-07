package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/olujimiAdebakin/ProtoGraph/account"
	"github.com/olujimiAdebakin/ProtoGraph/catalog"
	"github.com/olujimiAdebakin/ProtoGraph/order"
)

// Server is the root GraphQL resolver.
// It holds clients to your microservices (Account, Catalog, Order).
// gqlgen will use this struct to resolve queries, mutations, and nested fields.
type Server struct {
	accountClient *account.Client
	catalogClient *catalog.Client
	orderClient   *order.Client
}

// NewGraphQlServer initializes the Server struct.
// It creates connections to the three downstream services.
//
// If any client fails, function cleans up previously created clients
// to avoid resource leakage.
func NewGraphQlServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
	// connect to the account service
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}

	// connect to the catalog service
	catalogClient, err := catalog.NewClient(catalogUrl)
	if err != nil {
		accountClient.Close()
		return nil, err
	}

	// connect to the order service
	orderClient, err := order.NewClient(orderUrl)
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return nil, err
	}

	// return server instance containing all three clients
	return &Server{
		accountClient: accountClient,
		catalogClient: catalogClient,
		orderClient:   orderClient,
	}, nil
}

// Mutation returns the MutationResolver implementation.
// gqlgen calls this automatically whenever a mutation is executed.
//
// → This wires GraphQL "Mutation { ... }"
//   to your mutationResolver struct.
func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

// Query returns the QueryResolver implementation.
// gqlgen calls this anytime someone runs a query.
//
// → This wires GraphQL "Query { ... }"
//   to your queryResolver struct.
func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}

// Account returns the resolver for nested Account fields.
// Example: Account → orders
//
// gqlgen calls this when a field needs its own resolver.
func (s *Server) Account() AccountResolver {
	return &accountResolver{
		server: s,
	}
}

// ToExecutableSchema compiles your resolvers + schema into
// an executable GraphQL schema engine.
//
// This is what gets passed into the HTTP GraphQL handler.
//
// Without this function, your GraphQL server cannot run.
func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	cfg := Config{
		Resolvers: s, // tell gqlgen to use this Server as the resolver root
	}

	return NewExecutableSchema(cfg)
}
