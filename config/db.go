package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}

func Connect() *bun.DB {
	dsn := "postgres://postgres:postgres@localhost:5432/shorten_url?sslmode=disable"
	pgconn := sql.OpenDB(
		pgdriver.NewConnector(
			pgdriver.WithDSN(dsn),
			pgdriver.WithAddr("localhost:5432"),
			pgdriver.WithUser("postgres"),
			pgdriver.WithPassword(goDotEnvVariable("PASSWORD")),
			pgdriver.WithDatabase("links"),
		),
	)
	db := bun.NewDB(pgconn, pgdialect.New())

	return db
}
