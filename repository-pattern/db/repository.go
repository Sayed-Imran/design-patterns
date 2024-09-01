package db

import (
	"context"

	"github.com/sayed-imran/go-design-pattern/models"
)

type Repository interface {
	AddUser(ctx context.Context, u models.User) error
	AddMultipleUsers(ctx context.Context, usersTobeInserted ...models.User) error
	FindSingleUser(ctx context.Context, id int) (models.User, error)
	FindMultipleUsers(ctx context.Context, filter interface{}) ([]models.User, error)
	UpdateUser(ctx context.Context, id int, u models.User) error
	DeleteUser(ctx context.Context, id int) error
	DeleteAllUsers(ctx context.Context) error
	Disconnect(ctx context.Context) error
}
