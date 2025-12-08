package account

import (
	"context"
	"github.com/segmentio/ksuid"
	"github.com/golang.org/x/crypto/bcrypt"
	"fmt"
	"errors"
)

// Predefined errors for input validation
var (
    ErrInvalidName    = errors.New("account name cannot be empty")
    ErrInvalidEmail   = errors.New("email cannot be empty")
    ErrWeakPassword   = errors.New("password must be at least 8 characters")
)

// Service defines the business operations related to accounts.
type Service interface {
	// PostAccount creates a new account with the given name, email, and password.
	// Returns the created Account or an error if validation/storage fails.
	PostAccount(ctx context.Context, name, email, password string) (*Account, error)

	// GetAccount fetches an account by its unique ID.
	GetAccount(ctx context.Context, id string) (*Account, error)

	// ListAccounts returns a paginated list of accounts, skipping 'skip' and taking 'take' items.
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)

	// DeleteAccount removes an account by ID, returning the deleted account or an error.
	DeleteAccount(ctx context.Context, id string) (*Account, error)
}

// Account represents a user account with identifying and authentication data.
type Account struct {
	ID       string `json:"id"`       // Unique identifier for the account
	Name     string `json:"name"`     // User's name
	Email    string `json:"email"`    // User's email address
	Password string `json:"password"` // Hashed password do not expose in APIs
}

// accountService implements the Service interface by interacting with a repository.
type accountService struct {
	repository AccountRepository // Interface to the data layer for accounts
}

// newAccountService constructs a new Service implementation backed by a repository.
func newAccountService(repo AccountRepository) Service {
	return &accountService{repo}
}

// PostAccount validates input, hashes the password, and stores the new account.
func (s *accountService) PostAccount(ctx context.Context, name, email, password string) (*Account, error) {

	// Validate inputs - fail fast on invalid data
	if name == "" {
		return nil, ErrInvalidName
	}
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if len(password) < 8 {
		return nil, ErrWeakPassword
	}

	// Hash the password securely with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create the account struct with generated ID and hashed password
	a := &Account{
		Name:     name,
		ID:       ksuid.New().String(),
		Email:    email,
		Password: string(hashedPassword),
	}

	// Persist the account through the repository layer
	if err := s.repository.PutAccount(ctx, *a); err != nil {
		return nil, err
	}

	return a, nil
}

// GetAccount retrieves an account by ID via the repository.
func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	return s.repository.GetAccountByID(ctx, id)
}

// ListAccounts provides a paginated list of accounts.
// Caps the page size to 50 to prevent overloading.
func (s *accountService) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	if take > 50 || (skip == 0 && take == 0) {
		take = 100 // enforce a maximum page size
	}
	return s.repository.ListAccounts(ctx, skip, take)
}


func (s *accountService) DeleteAccount(ctx context.Context, id string) (*Account, error) {
	// 1. Get the account first
	acc, err := s.repository.GetAccountByID(ctx, id)
	if err != nil {
        return nil, err // could be not found or DB error
    }

    // 2. Delete the account
    err = s.repository.DeleteAccount(ctx, id)
    if err != nil {
        return nil, err
    }

    // 3. Return the deleted account
    return acc, nil
}