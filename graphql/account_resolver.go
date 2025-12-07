package main

import (
	"context"
	"time"
)

type accountResolver struct {
	server *Server
}

// CreatedAt implements AccountResolver.
func (a *accountResolver) CreatedAt(ctx context.Context, obj *Account) (*time.Time, error) {
	panic("unimplemented")
}

// Orders implements AccountResolver.
func (a *accountResolver) Orders(ctx context.Context, obj *Account) ([]*Order, error) {
	panic("unimplemented")
}

// UpdatedAt implements AccountResolver.
func (a *accountResolver) UpdatedAt(ctx context.Context, obj *Account) (*time.Time, error) {
	panic("unimplemented")
}
