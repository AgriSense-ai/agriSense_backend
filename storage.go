package main

import ("database/sql"
		_ "github.com/lib/pq"
)

type Storage interface {
   CreateAccount(*Account) error
   DeleteAccount(int) error
   UpdateAccount(*Account) error
   GetAccountByID(int) (*Account, error)
}

type postgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*postgressStore, error) {
	// connStr := "user=postgres password= dbname=bank sslmode=disable"
	// db, err := sql.Open(connStr)
	return nil, nil
}