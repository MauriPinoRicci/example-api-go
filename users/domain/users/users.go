package users

import (
	"errors"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type User struct {
	id       string
	name string
	email    string
	status   string
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
		id:       id,
		name: name,
		email:    email,
		status:   status,
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
	if u.id == "" {
		return errors.New("invalid ID")
	}
	if u.name == "" {
		return errors.New("invalid name")
	}
	if u.email == "" {
		return errors.New("invalid Email")
	}
	_, valid := AllStatus[u.status]
	if !valid {
		return errors.New("invalid Status")
	}
	return nil
}

func (s *User) normalize() {
	s.email = strings.ToLower(s.email)
	s.status = strings.ToUpper(s.status)
}

func (s *User) ID() string {
	return s.id
}

func (s *User) Name() string {
	return s.name
}

func (s *User) Email() string {
	return s.email
}

func (s *User) Status() string {
	return s.status
}
