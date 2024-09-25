package main

import (
	"fmt"
	"os"

	"github.com/ferdiebergado/fullstack-go/db"
	"github.com/ferdiebergado/fullstack-go/pkg/env"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	err := env.LoadEnv()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open .env file")
		os.Exit(1)
	}

	database := db.OpenDb()

	defer database.Close()

	router := NewApp(database)

	RunServer(router)
}
