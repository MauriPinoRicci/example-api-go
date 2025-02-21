package update_srv

import (
	"context"

	"github.com/MauriPinoRicci/example-api-go/users/application/shared"
	"github.com/MauriPinoRicci/example-api-go/users/domain/users"
)

type Service struct {
	repo users.Repository
}

type UpdateUserInput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func NewService(repo users.Repository) *Service {
	return &Service{repo}
}

func (s *Service) Update(ctx context.Context, input *UpdateUserInput) (*shared.UserOutput, error) {
    user, err := s.repo.GetByID(ctx, input.ID)
    if err != nil {
        return nil, err
    }

    if input.Name != "" {
        user.SetName(input.Name)
    }
    if input.Email != "" {
        user.SetEmail(input.Email)
    }
    if input.Status != "" {
        user.SetStatus(input.Status)
    }

    updatedUser, err := s.repo.Update(ctx, input.ID, user)
    if err != nil {
        return nil, err
    }

    return shared.BuildUserOutput(updatedUser), nil
}


