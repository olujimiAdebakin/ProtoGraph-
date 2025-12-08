package catalog

import "context"

type Repository interface {
    Close()
    PutProduct(ctx context.Context, p Product) error
    GetProductByID(ctx context.Context, id string) (*Product, error)
    ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error)
    DeleteProduct(ctx context.Context, id string) error
}

type postgresRepositry struct{
	db *spl.DB
}