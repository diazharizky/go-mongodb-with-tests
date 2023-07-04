package app

import (
	"context"

	"github.com/diazharizky/go-mongodb-with-tests/internal/models"
)

type IUserRepository interface {
	List(ctx context.Context) ([]models.User, error)
	Get(ctx context.Context, id string) (*models.User, error)
	Create(ctx context.Context, newUser models.User) (id *string, err error)
	Update(ctx context.Context, updateValues models.User) error
	Delete(ctx context.Context, id string) error
}
