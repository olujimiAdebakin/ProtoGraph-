package account

import (
    "context"
    "database/sql"
    _ "github.com/lib/pq"

)

// Account repository interface
type AccountRepository interface {
    Close()
    
    // Create or Update an account
    PutAccount(ctx context.Context, a Account) error

    // Fetch one account by ID
    GetAccountByID(ctx context.Context, id string) (*Account, error)

    // List accounts with pagination (skip = offset, take = limit)
    ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)

    // Delete account by ID
    DeleteAccount(ctx context.Context, id string) error
}

// // Product repository interface
// type ProductRepository interface {
//     Close()

//     PutProduct(ctx context.Context, p Product) error

//     GetProductByID(ctx context.Context, id string) (*Product, error)

//     ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error)

//     DeleteProduct(ctx context.Context, id string) error
// }

// // Order repository interface
// type OrderRepository interface {
//     Close()

//     PutOrder(ctx context.Context, o Order) error

//     GetOrderByID(ctx context.Context, id string) (*Order, error)

//     ListOrders(ctx context.Context, skip uint64, take uint64) ([]Order, error)

//     DeleteOrder(ctx context.Context, id string) error
// }

// postgresRepository implements the AccountRepository interface.
// It provides all DB operations for accounts using PostgreSQL.
type postgresRepositry struct{
	db *sql.DB
}

// NewPostgresRepository connects to PostgreSQL and returns a repository instance.
func NewPostgresRepositry(url string) (AccountRepository, error){
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil,err
	}

     // Verify connection
	err := db.Ping()
	if err != nil {
		return nil,err
	}

	return &postgresRepositry(db: db), nil
}



func (r *postgresRepositry) Close(){
	r.db.Close()
}

func (r *postgresRepositry) Ping() error{
	return r.db.Ping()
}

// PutAccount inserts or updates an account (UPSERT logic).
func (r *postgresRepositry) PutAccount(ctx context.Context, a Account) error{
	_, err := r.db.ExecContext(ctx, "INSERT INTO accounts (id, name, email) VALUES ($1, $2, $3) ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name, email = EXCLUDED.email", a.ID, a.Name, a.Email)
    return err
}

// GetAccountByID fetches a single account by ID.
// It returns (*Account, nil) if found.
// It returns (nil, nil) if no row exists.
// It returns (nil, error) for DB errors.
func (r *postgresRepositry) GetAcountByID(ctx context.Context, id string) (*Account, error){
    // Create the struct that will hold the scanned DB values
       acc := &Account{}

        // Query the database
	err := r.db.QueryRowContext(ctx, "SELECT id, name, email FROM accounts WHERE id = $1", id).Scan(&acc.ID, &acc.Name, &acc.Email)

   // If no row found, return nil instead of error
        if err == sql.ErrNoRows {
            return nil, nil // Account not found
        }

  
    // Actual error
    if err != nil {
        return nil, err
    }

    return acc, nil
}

// ListAccounts returns paginated accounts using LIMIT + OFFSET.
func (r *postgresRepositry) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error){
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email FROM accounts ORDER BY id OFFSER $1  LIMIT $2", skip, take)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    accounts := []Account{}

    for rows.Next() {
        a := Account{}
        err := rows.Scan(&a.ID, &a.Name, &a.Email)
        if err != nil {
            return nil, err
        }
        accounts = append(accounts, a)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return accounts, nil
}

// DeleteAccount removes an account by ID.
func (r *postgresRepositry) DeleteAccount(ctx context.Context, id string) error{
    _, err := r.db.ExecContext(ctx, "DELETE FROM accounts WHERE id = $1", id)
    return err
}