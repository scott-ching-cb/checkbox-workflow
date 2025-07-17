package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"workflow-code-test/api/pkg/db"
	"workflow-code-test/api/services/workflow"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// handleDatabaseMigration handles the database migration using golang-migrate package
func handleDatabaseMigration() error {
	// Connect to the database to initiate migrations
	migrationDatabase, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("failed to setup connect to database for migrations : %v", err)
	}

	// Create migration driver for Postgres database
	driver, err := postgres.WithInstance(migrationDatabase, &postgres.Config{
		MigrationsTable: "migrations",
		DatabaseName:    "workflow_engine",
	})
	if err != nil {
		return fmt.Errorf("failed to create driver for database migrations : %v", err)
	}

	migrationDatabaseInstance, err := migrate.NewWithDatabaseInstance(
		"file://./migrations/",
		"workflow_engine",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to instantiate new migration instance : %v", err)
	}

	defer func() {
		sourceErr, dbErr := migrationDatabaseInstance.Close()
		if sourceErr != nil {
			slog.Error("Error closing database connection", "error", sourceErr.Error())
			return
		}
		if dbErr != nil {
			slog.Error("Error closing database connection", "error", dbErr.Error())
			return
		}
	}()

	if migrationError := migrationDatabaseInstance.Up(); migrationError != nil {
		if errors.Is(migrationError, migrate.ErrNoChange) {
			return nil
		}
		if migrationRollbackError := migrationDatabaseInstance.Down(); migrationRollbackError != nil {
			slog.Error("Failed to rollback migration", "error", migrationRollbackError)
		}
		return fmt.Errorf("failed to execute migration : %v", migrationError)
	}
	return nil
}

func main() {
	// Configure structured logging
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slog.SetDefault(slog.New(logHandler))

	dbConfig := db.DefaultConfig()
	dbConfig.URI = os.Getenv("DATABASE_URL")
	if err := db.Connect(dbConfig); err != nil {
		slog.Error("Failed to connect to database", "error", err, "dbConfig", dbConfig.URI)
		return
	}
	defer db.Disconnect()

	if migrationError := handleDatabaseMigration(); migrationError != nil {
		slog.Error("Failed to execute database migration", "error", migrationError)
	}

	// setup router
	mainRouter := mux.NewRouter()

	apiRouter := mainRouter.PathPrefix("/api/v1").Subrouter()

	// Get a connection from the pool
	conn, err := db.GetPool().Acquire(context.Background())
	if err != nil {
		slog.Error("Failed to acquire database connection", "error", err)
		return
	}
	defer conn.Release()

	workflowService, err := workflow.NewService(conn.Conn())
	if err != nil {
		slog.Error("Failed to create workflow service", "error", err)
		return
	}

	workflowService.LoadRoutes(apiRouter, false)

	// Configure CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3003"}), // Frontend URL
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)(mainRouter)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: corsHandler,
	}

	// Channel to listen for errors coming from the server
	serverErrors := make(chan error, 1)

	// Start the server in a goroutine
	go func() {
		slog.Info("Starting server on :8080")
		serverErrors <- srv.ListenAndServe()
	}()

	// Channel to listen for an interrupt or terminate signal from the OS
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking select waiting for either a signal or an error
	select {
	case err := <-serverErrors:
		slog.Error("Server error", "error", err)

	case sig := <-shutdown:
		slog.Info("Shutdown signal received", "signal", sig)

		// Give outstanding requests 5 seconds to complete
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("Could not stop server gracefully", "error", err)
			srv.Close()
		}
	}
}
