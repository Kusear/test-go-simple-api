package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"test-go-simple-api/api"
	"test-go-simple-api/internal/services"

	"github.com/go-chi/chi/v5"
)

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type registrationResource struct {
	service *services.AccountService
}

func (rR registrationResource) Routes() chi.Router {
	router := chi.NewRouter()

	router.Post("/", rR.RegisterAccount)

	return router
}

func (rR registrationResource) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.InternalServerError(w)
		return
	}

	createdUser, err := rR.service.Create(ctx, user.Name, user.Username)
	if err != nil {
		api.InternalServerError(w)
		return
	}

	api.SuccessResponse(w, createdUser)
}
