package repositories

import (
	"context"
	"test-go-simple-api/internal/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (uR *UserRepository) Save(ctx context.Context, data *entities.User) (*entities.User, error) {

	result := gorm.WithResult()
	executor := gorm.G[entities.User](uR.Db, result)

	err := executor.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uR *UserRepository) Update(ctx context.Context, data *entities.User) (*entities.User, error) {
	executor := gorm.G[entities.User](uR.Db)
	_, err := executor.Updates(ctx, *data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uR *UserRepository) Delete(ctx context.Context, id int) error {
	userToDelete, err := uR.FindById(ctx, id)
	if err != nil {
		return err
	}

	executor := gorm.G[entities.User](uR.Db)
	_, err = executor.Updates(ctx, *userToDelete)
	if err != nil {
		return err
	}

	return nil
}

func (uR *UserRepository) FindById(ctx context.Context, id int) (*entities.User, error) {
	executor := gorm.G[entities.User](uR.Db)
	user, err := executor.Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
