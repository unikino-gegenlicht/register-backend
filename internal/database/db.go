package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"register-backend/internal/configuration"
	"register-backend/resources"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

var Pool *pgxpool.Pool

func init() {
	c := configuration.Config.Sub("database")
	if c == nil {
		slog.Error("database connection not configured")
		os.Exit(1)
	}
	if !c.IsSet("host") {
		slog.Error("database host not set")
		os.Exit(1)
	}

	if !c.IsSet("port") {
		c.Set("port", 5432)
	}

	if !c.IsSet("user") || !c.IsSet("password") {
		slog.Error("database credentials not set")
		os.Exit(1)
	}

	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%d sslmode=disable",
		c.GetString("user"), c.GetString("password"), c.GetString("host"), c.GetInt("port"))

	slog.Debug("initializing database connection", "connection_string", connectionString)
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		slog.Error("unable to create database pool", "error", err.Error())
		os.Exit(1)
	}
	if err := pool.Ping(context.Background()); err != nil {
		slog.Error("unable to connect to the database", "error", err.Error())
		os.Exit(1)
	}
	Pool = pool

	// execute migrations to the database against the database schema
	db := stdlib.OpenDBFromPool(Pool)
	goose.SetBaseFS(resources.Migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		slog.Error("unable to use posgres as goose dialect", "error", err.Error())
		os.Exit(1)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		slog.Error("unable to execute database migrations", "error", err.Error())
		os.Exit(2)
	}

}
