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

func (s *postgressStore) GetAccounts() ([]*Account, error) {
	query := `SELECT * FROM account;`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		acc, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, acc)
	}
	return accounts, nil
}

func (s *postgressStore) GetAccountByID(id int) (*Account, error) {
	rows, err := s.db.Query(`SELECT * FROM account WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account %d not found", id)
}

func (s *postgressStore) UpdateAccount(*Account) error {
	return nil
}

func (s *postgressStore) DeleteAccount(id int) error {
	query := `DELETE FROM account WHERE id = $1;`
	_, err := s.db.Exec(query, id)
	return err
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	acc := new(Account)
	err := rows.Scan(
		&acc.ID,
		&acc.FirstName,
		&acc.LastName,
		&acc.Number,
		&acc.Balance,
		&acc.CreatedAt)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
