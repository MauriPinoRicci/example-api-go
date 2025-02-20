package shared

import "github.com/MauriPinoRicci/example-api-go/users/domain/users"

type UserOutput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func BuildUserOutput(s *users.User) *UserOutput {
	return &UserOutput{
		ID:     s.ID(),
		Name:   s.Name(),
		Email:  s.Email(),
		Status: s.Status(),
	}
}
