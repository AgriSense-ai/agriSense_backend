package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type postgressStore struct {
	db *sql.DB
}

// GetAccounts implements Storage.
func (s *postgressStore) GetAccounts() ([]*Account, error) {
	panic("unimplemented")
}

func NewPostgressStore() (*postgressStore, error) {
	connStr := "user=root password=ConradKash dbname=agrisense sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &postgressStore{db: db}, nil
}

func (s *postgressStore) Init() error {
	return s.createAccountTable()
}

func (s *postgressStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
				id SERIAL PRIMARY KEY,
				first_name VARCHAR(50), 
				last_name VARCHAR(50),
				number SERIAL,
				balance INT,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
				);`
	_, err := s.db.Exec(query)
	return err
}

func (s *postgressStore) CreateAccount(acc *Account) error {
	query := `INSERT INTO account (first_name, last_name, number, balance, created_at) 
			VALUES ($1, $2, $3, $4, $5);`
	resp, err := s.db.Exec(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt,
	)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return nil
}

func (s *postgressStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}

func (s *postgressStore) UpdateAccount(*Account) error {
	return nil
}

func (s *postgressStore) DeleteAccount(id int) error {
	return nil
}
