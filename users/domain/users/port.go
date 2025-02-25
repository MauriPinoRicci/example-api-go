package users

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	Delete(ctx context.Context, id string) error
}
