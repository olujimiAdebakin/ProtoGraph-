package order

import "context"

type Repository interface {
    Close()
    PutOrder(ctx context.Context, o Order) error
    GetOrderByID(ctx context.Context, id string) (*Order, error)
    ListOrders(ctx context.Context, skip uint64, take uint64) ([]Order, error)
    DeleteOrder(ctx context.Context, id string) error
}

type postgresRepositry struct{
	db *spl.DB
}