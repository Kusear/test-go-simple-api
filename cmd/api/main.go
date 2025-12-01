package main

import (
	"context"
	"database/sql"
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
)

func main() {
	var router *chi.Mux = chi.NewRouter()

	dbPool, err := initDatabase()
	if err != nil {
		fmt.Printf("Database error: %v\n", err)
		return
	}

	handlers.RegisterHandlers(router, handlers.Infrastructure{
		DbConnection: dbPool,
	})

	fmt.Printf("Database inited %v\n", err)

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Error starting server:", err)
			return
		}
	}()

	fmt.Println("Server started at 8080 port...")

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

func initDatabase() (*sql.DB, error) {
	// TODO add env variables and take it from there
	dbConn := database.DatabaseConnector{
		Host:     "localhost",
		Port:     5432,
		Username: "test",
		Password: "test",
		DBName:   "go_todo_lists",
	}

	dbPool, err := dbConn.Connect()
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}
