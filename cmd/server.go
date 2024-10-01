package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/ferdiebergado/fullstack-go/internal/config"
	"github.com/ferdiebergado/fullstack-go/internal/db"
	"github.com/ferdiebergado/fullstack-go/pkg/stdout"
)

func RunServer(ctx context.Context, w io.Writer, args []string, port string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	conn := db.OpenDb()
	defer conn.Close()

	queries := db.New(conn)

	router := NewApp(conn, queries)

	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  config.ServerReadTimeout,
		WriteTimeout: config.ServerWriteTimeout,
		IdleTimeout:  config.ServerIdleTimeout,
	}

	go func() {
		fmt.Printf("%sHTTP Server listening on %s...%s\n", stdout.Cyan, httpServer.Addr, stdout.Reset)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()

	return nil
}
