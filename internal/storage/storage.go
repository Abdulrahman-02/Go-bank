package storage

import (
	"database/sql"

	"github.com/Abdulrahman-02/Go-bank/internal/types"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*types.Account) error
	DeleteAccount(int) error
	UpdateAccount(*types.Account) error
	GetAccount(int) (*types.Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres-1 dbname=bank password=password sslmode=disable"
	db, err := sql.Open("postgres-1", connStr)
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) CreateAccount(*types.Account) error {
	return nil
}

func (s *PostgresStorage) DeleteAccount(int) error {
	return nil
}

func (s *PostgresStorage) UpdateAccount(*types.Account) error {
	return nil
}

func (s *PostgresStorage) GetAccount(int) (*types.Account, error) {
	return nil, nil
}
