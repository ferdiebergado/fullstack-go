package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ferdiebergado/fullstack-go/config"
	"github.com/ferdiebergado/fullstack-go/pkg/env"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
)

// GracefulShutdown gracefully shuts down the server when receiving termination signals.
func GracefulShutdown(srv *http.Server, wg *sync.WaitGroup) {
	defer wg.Done()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Println("Server exited properly")
}

func RunServer(router *myhttp.Router) {
	port := env.GetEnv("APP_PORT")

	// Create the server.
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  config.ServerReadTimeout,
		WriteTimeout: config.ServerWriteTimeout,
		IdleTimeout:  config.ServerIdleTimeout,
	}

	// Run the server with graceful shutdown.
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go GracefulShutdown(server, wg)

	fmt.Printf("Server started on port %s...\n", port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %v", err)
	}

	wg.Wait()
}
