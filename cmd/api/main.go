package main

import (
	"context"
	"fmt"
	"github.com/sunil-pateel/personal-website/internal/server"
	"log"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {
	const PORT int = 5001
	done := make(chan bool, 1)

	server := server.NewServer(PORT)
	domain := fmt.Sprintf("http://localhost:%d/", PORT)
	slog.Info("starting server", "domain", domain)

	go gracefulShutdown(server, done)
	server.ListenAndServe()

	<-done

    log.Println("Graceful shutdown complete")

}
