package account

import ()


type Service struct {
	PostAccount()
	GetAccount()
	ListAccounts()
	DeleteAccount()
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type accountService struct {
	repository Repository
}