package repositories

import "context"

type Repository[T interface{}] interface {
	Save(ctx context.Context, data *T) (*T, error)
	Update(ctx context.Context, data *T) (*T, error)
	Delete(ctx context.Context, id int) error

	// find(ctx context.Context, filter )
	FindById(ctx context.Context, id int) (*T, error)
}
