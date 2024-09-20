package main

import (
	"github.com/ferdiebergado/fullstack-go/internal/db"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	database := db.OpenDb()

	defer database.Close()

	router := NewApp()

	RunServer(router)
}
