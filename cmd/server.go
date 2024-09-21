package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/ferdiebergado/fullstack-go/internal/config"
	"github.com/ferdiebergado/fullstack-go/pkg/env"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
)

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
	go myhttp.GracefulShutdown(server, wg)

	fmt.Printf("Server started on port %s...\n", port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %v", err)
	}

	wg.Wait()
}
