package users_dynamo

import "github.com/MauriPinoRicci/example-api-go/users/domain/users"

type UserMsg struct {
	ID     string `dynamodbav:"id,omitempty" json:"id,omitempty"`
	Name   string `dynamodbav:"name,omitempty" json:"name,omitempty"`
	Email  string `dynamodbav:"email,omitempty" json:"email,omitempty"`
	Status string `dynamodbav:"status,omitempty" json:"status,omitempty"`
}

func BuildUserMsg(s *users.User) *UserMsg {
	resp := &UserMsg{
		ID:     s.ID(),
		Name:   s.Name(),
		Email:  s.Email(),
		Status: s.Status(),
	}

	return resp
}

func (s *UserMsg) ToUser() (*users.User, error) {
	return users.NewUsers(
		s.ID,
		s.Name,
		s.Email,
		s.Status,
	)
}
