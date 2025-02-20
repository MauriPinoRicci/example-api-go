package delete_srv

import (
	"context"

	"github.com/MauriPinoRicci/example-api-go/users/domain/users"
)

type Service struct {
	repo users.Repository
}

type DeleteUserInput struct {
	ID string `json:"id"`
}

func NewService(repo users.Repository) *Service {
	return &Service{repo}
}

func (s *Service) Execute(ctx context.Context, input *DeleteUserInput) (string, error) {
	err := s.repo.Delete(ctx, input.ID)
	if err != nil {
		return "", err
	}
	return "Usuario eliminado correctamente", nil
}
