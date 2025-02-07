package users

import (
	"errors"
	"strings"
)

type User struct {
	id       string
	username string
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

func NewUsers(id, username, email, status string) (*User, error) {
	user := &User{
		id:       id,
		username: username,
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

func (u *User) validate() error {
	if u.id == "" {
		return errors.New("invalid ID")
	}
	if u.username == "" {
		return errors.New("invalid Username")
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

func (s *User) Username() string {
	return s.username
}

func (s *User) Email() string {
	return s.email
}

func (s *User) Status() string {
	return s.status
}
