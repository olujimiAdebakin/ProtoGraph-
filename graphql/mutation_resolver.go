package main

import "context"

type mutationResolver struct {
	server *Server
}

// CreateAccount implements MutationResolver.
func (m *mutationResolver) CreateAccount(ctx context.Context, input AccountInput) (*Account, error) {
	panic("unimplemented")
}

// CreateOrder implements MutationResolver.
func (m *mutationResolver) CreateOrder(ctx context.Context, input OrderInput) (*Order, error) {
	panic("unimplemented")
}

// CreateProduct implements MutationResolver.
func (m *mutationResolver) CreateProduct(ctx context.Context, input ProductInput) (*Product, error) {
	panic("unimplemented")
}

// DeleteAccount implements MutationResolver.
func (m *mutationResolver) DeleteAccount(ctx context.Context, id string) (bool, error) {
	panic("unimplemented")
}

// DeleteOrder implements MutationResolver.
func (m *mutationResolver) DeleteOrder(ctx context.Context, id string) (bool, error) {
	panic("unimplemented")
}

// DeleteProduct implements MutationResolver.
func (m *mutationResolver) DeleteProduct(ctx context.Context, id string) (bool, error) {
	panic("unimplemented")
}

// UpdateAccount implements MutationResolver.
func (m *mutationResolver) UpdateAccount(ctx context.Context, id string, input AccountInput) (*Account, error) {
	panic("unimplemented")
}

// UpdateOrder implements MutationResolver.
func (m *mutationResolver) UpdateOrder(ctx context.Context, id string, input OrderInput) (*Order, error) {
	panic("unimplemented")
}

// UpdateProduct implements MutationResolver.
func (m *mutationResolver) UpdateProduct(ctx context.Context, id string, input ProductInput) (*Product, error) {
	panic("unimplemented")
}
