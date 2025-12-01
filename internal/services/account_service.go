package services

import (
	"context"
	"database/sql"
	"fmt"
)

type AccountService struct {
	dbConnection *sql.DB
}

func (aS AccountService) Create(ctx context.Context) error {
	fmt.Println("TODO account service Create call")
	return nil
}

func (aS AccountService) Update(ctx context.Context) error {
	fmt.Println("TODO account service Update call")

	return nil
}

func (aS AccountService) GetAccountInfo(ctx context.Context) error {
	fmt.Println("TODO account service GetAccountInfo call")

	return nil
}

func InitAccountService(db *sql.DB) *AccountService {
	return &AccountService{
		dbConnection: db,
	}
}
