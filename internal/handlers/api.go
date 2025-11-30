package handlers

import (
	"test-go-simple-api/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterHandlers(router *chi.Mux) {

	// router.Use(middleware.AuthMiddleware)
	router.Use(middleware.LoggingMiddleware)

	router.Route("/register", func(r chi.Router) {
		r.Post("/", PostCreateAccount)
	})
	router.Route("/account", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/", GetAccount)
	})
}
