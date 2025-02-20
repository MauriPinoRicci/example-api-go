package get_srv

import (
	"context"

	"github.com/MauriPinoRicci/example-api-go/users/application/shared"
	"github.com/MauriPinoRicci/example-api-go/users/domain/users"
)

type Service struct {
	repo users.Repository
}

type GetUserInput struct {
	ID string `json:"id"`
}

func NewService(repo users.Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetByID(ctx context.Context, input *GetUserInput) (*shared.UserOutput, error) {
	user, err := s.repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return shared.BuildUserOutput(user), nil
}
