package main

import "context"

type queryResolver struct {
	server *Server
}

// GetAccount implements QueryResolver.
func (q *queryResolver) GetAccount(ctx context.Context, id string) (*Account, error) {
	panic("unimplemented")
}

// GetProduct implements QueryResolver.
func (q *queryResolver) GetProduct(ctx context.Context, id string) (*Product, error) {
	panic("unimplemented")
}

// ListAccounts implements QueryResolver.
func (q *queryResolver) ListAccounts(ctx context.Context, pagination *PaginationInput) ([]*Account, error) {
	panic("unimplemented")
}

// ListProducts implements QueryResolver.
func (q *queryResolver) ListProducts(ctx context.Context, pagination *PaginationInput) ([]*Product, error) {
	panic("unimplemented")
}
