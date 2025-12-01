package handlers

import (
	"database/sql"
	"test-go-simple-api/internal/middleware"
	"test-go-simple-api/internal/services"

	"github.com/go-chi/chi/v5"
)

type Infrastructure struct {
	DbConnection *sql.DB
}

func RegisterHandlers(router *chi.Mux, infra Infrastructure) {

	router.Use(middleware.LoggingMiddleware)

	accountService := services.InitAccountService(infra.DbConnection)

	router.Mount("/account", accountResource{
		service: accountService,
	}.Routes())

	router.Mount("/register", registrationResource{}.Routes())
	// router.Route("/register", func(r chi.Router) {
	// 	r.Post("/", PostCreateAccount)
	// })
}
