package store_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"testing"
	"time"
	"workflow-code-test/api/pkg/db"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	DB           *sql.DB
	DBConnection *pgx.Conn
)

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "16",
		Env: []string{
			"POSTGRES_PASSWORD=workflow123",
			"POSTGRES_USER=workflow",
			"POSTGRES_DB=workflow_engine_test",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://workflow:workflow123@%s/workflow_engine_test?sslmode=disable", hostAndPort)
	log.Println("Connecting to database on url: ", databaseUrl)
	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		DB, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return DB.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Create migration driver for Postgres database
	driver, err := postgres.WithInstance(DB, &postgres.Config{
		MigrationsTable: "migrations",
		DatabaseName:    "workflow_engine",
	})
	if err != nil {
		log.Fatalf("failed to create driver for database migrations : %v", err)
		return
	}

	migrationDatabaseInstance, err := migrate.NewWithDatabaseInstance(
		"file://../../../migrations/",
		"workflow_engine",
		driver,
	)
	if err != nil {
		log.Fatalf("failed to instantiate new migration instance : %v", err)
		return
	}

	defer func() {
		sourceErr, dbErr := migrationDatabaseInstance.Close()
		if sourceErr != nil {
			log.Fatalf("Error closing database connection: %v", sourceErr)
			return
		}
		if dbErr != nil {
			log.Fatalf("Error closing database connection: %v", dbErr)
		}
	}()

	if migrationError := migrationDatabaseInstance.Up(); migrationError != nil {
		log.Fatalf("Error running migrations: %v", migrationError)
		return
	}

	dbConfig := db.DefaultConfig()
	dbConfig.URI = databaseUrl
	if err = db.Connect(dbConfig); err != nil {
		log.Fatalf("Failed to connect to database")
		return
	}

	conn, err := db.GetPool().Acquire(context.Background())
	if err != nil {
		slog.Error("Failed to acquire database connection", "error", err)
		return
	}
	defer conn.Release()
	DBConnection = conn.Conn()

	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()
	m.Run()
}
