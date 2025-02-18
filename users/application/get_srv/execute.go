package get_srv

import (
	"context"

	"github.com/MauriPinoRicci/example-api-go/users/domain/users"
)

type Service struct {
	repo users.Repository
}

type GetUserInput struct {
	ID string `json:"id"`
}

type GetUserOutput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func NewService(repo users.Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetByID(ctx context.Context, input *GetUserInput) (*GetUserOutput, error) {
	user, err := s.repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &GetUserOutput{
		ID:     user.ID(),
		Name:   user.Name(),
		Email:  user.Email(),
		Status: user.Status(),
	}, nil

}
