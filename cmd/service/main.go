package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"

	api "github.com/ryanrolds/go_web_service/internal/api"
	"github.com/ryanrolds/go_web_service/internal/persistence"
)

func main() {
	// Setup logger/Logrus
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	// Get port to start HTTP server on
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT env var not provided, exiting")
	}
	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		log.Fatalf("PORT env var not a number (%v), exiting", port)
	}

	// DB connection string
	dbUrl := os.Getenv("POSTGRES_URL")
	if dbUrl == "" {
		log.Fatal("POSTGRES_URL env var not provided, exiting")
	}

	// Create DB connection
	conn, err := persistence.NewDBConnection(dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to DB (%v), exiting", err)
	}

	// Use wait group to track when Go routines are done and we can exit
	wg := &sync.WaitGroup{}

	// Use root context to notify Go routines to stop
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	log.Infof("Starting Services Domain API on port %d", port)
	api.StartServer(ctx, wg, port, conn)

	shutdownSignals := make(chan os.Signal, 1)
	signal.Notify(shutdownSignals, syscall.SIGINT, syscall.SIGTERM)

	// Wait for shutdown signal in background
	go func() {
		<-shutdownSignals
		log.Info("Received shutdown signal, beggining shutdown")

		// Cancel context, which notifies all Go routines to stop
		cancel()
	}()

	// Block main thread until all Go routines are done
	wg.Wait()

	// Now that everything is shutdown, close the DB connection
	err = conn.Close()
	if err != nil {
		log.Errorf("Failed to close DB connection (%v)", err)
	}

	log.Info("Gracefully shutdown")
}
