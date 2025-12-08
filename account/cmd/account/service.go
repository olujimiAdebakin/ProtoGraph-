package account

import (
	"context"
	"github.com/segmentio/ksuid"
	"github.com/golang.org/x/crypto/bcrypt"
	"fmt"
	"errors"
)

var (
    ErrInvalidName    = errors.New("account name cannot be empty")
    ErrInvalidEmail   = errors.New("email cannot be empty")
    ErrWeakPassword   = errors.New("password must be at least 8 characters")
)


type Service interface {
	// Create a new account by name, return the new Account or error.
	PostAccount(ctx context.Context, name, email, password string) (*Account, error)
	// Fetch an account by its unique ID.
	GetAccount(ctx context.Context, id string) (*Account, error)
	// List accounts with pagination (skip = offset, take = limit).
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
	// Delete an account by its unique ID, return the deleted Account or error.
	DeleteAccount(ctx context.Context, id string) (*Account, error)
}

// Account represents a user account in the system.
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"` 
	Password string `json:"password"`
}

type accountService struct {
	repository AccountRepository
}

func newAccountService(repo AccountRepository) Service {
	return &accountService{repo}
}



func (s *accountService) PostAccount(ctx context.Context, name, email, password string) (*Account, error) {

	    if name == "" {
        return nil, ErrInvalidName
    }
    if email == "" {
        return nil, ErrInvalidEmail
    }
    if len(password) < 8 {
        return nil, ErrWeakPassword
    }

    // Hash the password securely (bcrypt example)
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, fmt.Errorf("failed to hash password: %w", err)
    }
  
	// Create a new Account instance
	a := &Account{
		Name: name,
		ID: ksuid.New().String(),
		Email: email,
		Password: string(hashedPassword),
	}
	if err := s.repository.PutAccount(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	s.repository.GetAccountByID()
}

func (s *accountService) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	s.repository.ListAccounts()
}

func (s *accountService) DeleteAccount(ctx context.Context, id string) (*Account, error) {
	s.repository.DeleteAccount()
}