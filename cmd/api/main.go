package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test-go-simple-api/internal/database"
	"test-go-simple-api/internal/handlers"
	"time"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

const APP_PORT int = 3000

func main() {
	var router *chi.Mux = chi.NewRouter()

	dbPool, err := initDatabase()
	if err != nil {
		fmt.Printf("Database error: %v\n", err)
		return
	}

	err = runMigrations(dbPool)
	if err != nil {
		fmt.Printf("Database migration error: %v\n", err)
		return
	}

	handlers.RegisterHandlers(router, handlers.Infrastructure{
		DbConnection: dbPool,
	})

	fmt.Printf("Database initialized\n")

	srv := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", APP_PORT),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Error starting server:", err)
			return
		}
	}()

	fmt.Printf("Server started at %d port...\n", APP_PORT)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Error shutting down server:", err)
		return
	}
}

func initDatabase() (*gorm.DB, error) {
	// TODO add env variables and take it from there
	dbConn := database.DatabaseConnector{
		Host:     "localhost",
		Port:     5432,
		Username: "test",
		Password: "test",
		DBName:   "go_todo_lists",
	}

	dbPool, err := dbConn.Connect2()
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}

func runMigrations(db *gorm.DB) error {
	err := database.Up(db)
	if err != nil {
		fmt.Printf("Error running migrations: %v\n", err)
		return err
	}

	fmt.Println("Migrations run successfully")
	return nil
}
