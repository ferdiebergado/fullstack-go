package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
)

func RunServer(router *myhttp.Router)  {
	port, ok := os.LookupEnv("APP_PORT")

	if !ok {
		fmt.Fprintln(os.Stderr, "APP_PORT not set")
		os.Exit(1)
	}

	// Create the server.
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	// Run the server with graceful shutdown.
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go myhttp.GracefulShutdown(server, wg)

	fmt.Printf("Server started on port %s...\n", port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %v", err)
	}

	wg.Wait()
}