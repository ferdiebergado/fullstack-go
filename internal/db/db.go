package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

func OpenDb() *sql.DB {
	fmt.Println("Connecting to the database...")

	// dsn, ok := os.LookupEnv("DATABASE_URL")

	// if !ok {
	// 	fmt.Fprintln(os.Stderr, "DATABASE_URL not set")
	// 	os.Exit(1)
	// }

	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))

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
