package handlers

import (
	"net/http"
	"test-go-simple-api/internal/middleware"
	"test-go-simple-api/internal/services"

	"github.com/go-chi/chi/v5"
)

type accountResource struct {
	service *services.AccountService
}

func (aR accountResource) Routes() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.AuthMiddleware)
	router.Get("/", aR.GetAccount)

	return router
}

func (aR accountResource) GetAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	aR.service.GetAccountInfo(ctx, 1)
}
