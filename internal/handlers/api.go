package handlers

import (
	"test-go-simple-api/internal/repositories"
	"test-go-simple-api/internal/services"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

type Infrastructure struct {
	DbConnection *gorm.DB
}

func RegisterHandlers(router *chi.Mux, infra Infrastructure) {

	router.Use(chi_middleware.Logger)

	userRepository := repositories.UserRepository{
		Db: infra.DbConnection,
	}
	accountService := services.InitAccountService(&userRepository)

	router.Mount("/account", accountResource{
		service: accountService,
	}.Routes())

	router.Mount("/register", registrationResource{
		service: accountService,
	}.Routes())
}
