package users

import (
	"errors"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type User struct {
	Id     string
	Name   string
	Email  string
	Status string
}

const (
	StatusActive    string = "ACTIVE"
	StatusInactive  string = "INACTIVE"
	StatusDeleted   string = "DELETED"
	StatusSuspended string = "SUSPENDED"
)

var AllStatus map[string]struct{} = map[string]struct{}{
	StatusActive:    {},
	StatusInactive:  {},
	StatusDeleted:   {},
	StatusSuspended: {},
}

func NewUsers(id, name, email, status string) (*User, error) {
	user := &User{
		Id:     id,
		Name:   name,
		Email:  email,
		Status: status,
	}

	user.normalize()

	err := user.validate()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(name string) (*User, error) {
	id, err := gonanoid.New(10)
	if err != nil {
		return nil, err
	}
	return NewUsers(id, name, "test@gmail.com", StatusActive)
}

func (u *User) validate() error {
	if u.Id == "" {
		return errors.New("invalid ID")
	}
	if u.Name == "" {
		return errors.New("invalid name")
	}
	if u.Email == "" {
		return errors.New("invalid Email")
	}
	_, valid := AllStatus[u.Status]
	if !valid {
		return errors.New("invalid Status")
	}
	return nil
}

func (s *User) normalize() {
	s.Email = strings.ToLower(s.Email)
	s.Status = strings.ToUpper(s.Status)
}

func (s *User) GetID() string {
	return s.Id
}

func (s *User) GetName() string {
	return s.Name
}

func (s *User) GetEmail() string {
	return s.Email
}

func (s *User) GetStatus() string {
	return s.Status
}
