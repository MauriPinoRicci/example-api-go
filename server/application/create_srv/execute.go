package create_srv

import (
	"context"

	"github.com/MauriPinoRicci/example-api-go/server/domain/users"
)

type  Service struct{
	repo users.Repository 
}

func NewService(repo users.Repository) *Service {
	return &Service{repo}
}

func (s *Service) CreateUser(ctx context.Context, user *users.User) error {
	return s.repo.Create(ctx, user)
}

func (s *Service) GetUserByID(ctx context.Context, id string) (*users.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s*Service) UpdateUser(ctx context.Context, user *users.User) error{
	return s.repo.UpdateUser(ctx, user)
}

func (s*Service) DeleteUser(ctx context.Context, id string) error{
	return s.repo.DeleteUser(ctx, id)
}
