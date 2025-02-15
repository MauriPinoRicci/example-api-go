package create_srv

import (
	"context"
	"errors"

	"github.com/MauriPinoRicci/example-api-go/server/domain/users"
)

type Service struct {
	repo users.Repository
}

type CreateUserInput struct {
	Name string `json:"name"`
}

type CreateUserOutput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
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

func (s *Service) CreateUser(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error) {

	user, err := users.CreateUser(input.Name)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		ID:     user.Id,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}, nil
}

func (s *Service) GetUserByID(ctx context.Context, Id string) (*GetUserOutput, error) {
	user, err := s.repo.GetByID(ctx, Id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return &GetUserOutput{
		ID:     user.Id,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}, nil
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
