package user_in_mem

import (
	"context"
	"fmt"

	"github.com/MauriPinoRicci/example-api-go/users/domain/users"
)

type UsersInMem struct {
	users map[string]*users.User
}

func NewUsersInMem() *UsersInMem {
	return &UsersInMem{
		users: make(map[string]*users.User),
	}
}

var _ users.Repository = (*UsersInMem)(nil) // implement interface

func (s *UsersInMem) Save(ctx context.Context, entity *users.User) error {
	s.users[entity.ID()] = entity
	return nil
}

func (s *UsersInMem) GetByID(ctx context.Context, id string) (*users.User, error) {
	user, exists := s.users[id]
	if !exists {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}

func (s *UsersInMem) Delete(ctx context.Context, id string) error {
	_, exists := s.users[id]
	if !exists {
		return nil
	}

	delete(s.users, id)
	return nil
}
