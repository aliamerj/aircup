package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/aliamerj/aircup/api/server"
	"github.com/aliamerj/aircup/config"
)

func Run(cfg config.Config) error {
	dburl, err := createDBURL()
	if err != nil {
		return err
	}

	appServer, httpServer, err := server.NewServer(cfg, dburl)
	if err != nil {
		return err
	}

	done := make(chan bool, 1)
	go gracefulShutdown(httpServer, appServer, done)

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	<-done
	log.Println("Graceful shutdown complete.")
	return nil
}

func gracefulShutdown(httpServer *http.Server, appServer *server.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("server forced to shutdown with error: %v", err)
	}

	if err := appServer.Close(); err != nil {
		log.Printf("failed to close resources: %v", err)
	}

	log.Println("server exiting")
	done <- true
}

func createDBURL() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve home dir: %w", err)
	}

	return filepath.Join(home, ".local", "share", "Aircup", "aircup.db"), nil
}
