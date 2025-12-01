package handlers

import (
	"net/http"
	"test-go-simple-api/internal/services"

	"github.com/go-chi/chi/v5"
)

type registrationResource struct {
	service *services.AccountService
}

func (rR registrationResource) Routes() chi.Router {
	router := chi.NewRouter()

	router.Post("/", rR.RegisterAccount)

	return router
}

func (rR registrationResource) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rR.service.Create(ctx)
}
