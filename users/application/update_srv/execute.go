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
	ID     string
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
	user.UpdateUser(input.Name, input.Email, input.Status)
	err = s.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	return shared.BuildUserOutput(user), nil
}
