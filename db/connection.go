package db

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/smnschmnck/roundest-go-htmx/db/queries"
	"github.com/smnschmnck/roundest-go-htmx/db/seed"
)

var DB *queries.Queries

func InitDb() error {
	connectionString := os.Getenv("DATABASE_URL")
	if len(connectionString) <= 0 {
		return errors.New("no db connection string")
	}

	conn, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return err
	}

	DB = queries.New(conn)

	seed.Seed(DB)

	return nil
}
