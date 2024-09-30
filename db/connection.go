package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/ferdiebergado/fullstack-go/config"
	"github.com/ferdiebergado/fullstack-go/pkg/stdout"
)

const (
	connStr = "postgres://%s:%s@%s:%s/%s?sslmode=disable"
)

func OpenDb() *sql.DB {
	fmt.Print("Connecting to the database... ")

	driver := config.Db.Driver
	user := config.Db.User
	password := config.Db.Password
	host := config.Db.Host
	port := config.Db.Port
	database := config.Db.Name

	dsn := fmt.Sprintf(connStr, user, password, host, port, database)

	db, err := sql.Open(driver, dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = db.PingContext(context.Background())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Ping database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(stdout.Green + "connected." + stdout.Reset)

	return db
}
