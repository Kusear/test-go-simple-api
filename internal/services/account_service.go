package services

import (
	"context"
	"fmt"
	"test-go-simple-api/internal/entities"
	"test-go-simple-api/internal/repositories"
)

type AccountService struct {
	Repository repositories.Repository[entities.User]
}

func (aS *AccountService) Create(ctx context.Context, name string, username string) (*entities.User, error) {

	user, err := aS.Repository.Save(ctx, &entities.User{
		Name:     name,
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (aS *AccountService) Update(ctx context.Context, id int) error {
	fmt.Println("TODO account service Update call")
	return nil
}

func (aS *AccountService) GetAccountInfo(ctx context.Context, id int) (*entities.User, error) {

	user, err := aS.Repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func InitAccountService(repo repositories.Repository[entities.User]) *AccountService {
	return &AccountService{
		Repository: repo,
	}
}
