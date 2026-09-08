package user

import (
	"fmt"
	"strings"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateUser(name string) error {
	name = strings.TrimSpace(name)

	if len(name) < 2 {
		return ErrInvalidName
	}

	if err := s.repo.Save(name); err != nil {
		return fmt.Errorf("create user %q: %w", name, err)
	}

	return nil
}
