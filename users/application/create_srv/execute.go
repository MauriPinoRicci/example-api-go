package create_srv

import (
	"context"

	"github.com/MauriPinoRicci/example-api-go/users/application/shared"
	"github.com/MauriPinoRicci/example-api-go/users/domain/users"
)

type Service struct {
	repo users.Repository
}

type CreateUserInput struct {
	Name string `json:"name"`
}


func NewService(repo users.Repository) *Service {
	return &Service{repo}
}

func (s *Service) Execute(ctx context.Context, input *CreateUserInput) (*shared.UserOutput, error) {

	user, err := users.CreateUser(input.Name)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return shared.BuildUserOutput(user), nil 
}