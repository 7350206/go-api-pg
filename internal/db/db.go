package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // pg driver ??
)

type Database struct {
	Client *sqlx.DB
}

// constructor
func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
	)

	// Connect to a database and verify with a ping.
	dbConn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return &Database{}, fmt.Errorf("error when connect db: %w", err)
	}

	return &Database{Client: dbConn}, nil
}

// helpchecker
func (d Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
