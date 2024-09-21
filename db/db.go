package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/ferdiebergado/fullstack-go/pkg/env"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func OpenDb() *sql.DB {
	const dbUrlStr = "postgres://%s:%s@%s:%s/%s?sslmode=disable"

	fmt.Println("Connecting to the database...")

	dbUser := env.GetEnv("DB_USER")
	dbPass := env.GetEnv("DB_PASS")
	dbHost := env.GetEnv("DB_HOST")
	dbPort := env.GetEnv("DB_PORT")
	dbName := env.GetEnv("DB_NAME")
	dbDriver := env.GetEnv("DB_DRIVER")

	dsn := fmt.Sprintf(dbUrlStr, dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open(dbDriver, dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = db.PingContext(context.Background())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Ping database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to the database.")

	return db
}
