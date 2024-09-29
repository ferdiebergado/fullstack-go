package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/ferdiebergado/fullstack-go/pkg/env"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	dbDriver = "pgx"
	connStr  = "postgres://%s:%s@%s:%s/%s?sslmode=disable"
)

func OpenDb() *sql.DB {
	fmt.Print("Connecting to the database... ")

	user := env.GetEnv("DB_USER")
	password := env.GetEnv("DB_PASS")
	host := env.GetEnv("DB_HOST")
	port := env.GetEnv("DB_PORT")
	database := env.GetEnv("DB_NAME")

	dsn := fmt.Sprintf(connStr, user, password, host, port, database)

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

	fmt.Println("connected.")

	return db
}
