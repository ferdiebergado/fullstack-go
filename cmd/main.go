package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ferdiebergado/fullstack-go/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	if err := RunServer(context.Background(), os.Stdout, os.Args, config.App.Port); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
