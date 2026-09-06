package main

import (
	"errors"
	"fmt"
)

type UserRepository interface {
	Save(name string) error
}

type MemoryUserRepository struct{}

type UserService struct {
	repo UserRepository
}

var ErrUserAlreadyExists = errors.New("user already exists")

func main() {
	repo := MemoryUserRepository{}

	service := CreateService(repo)

	err := service.CreateUser("Alex")

	if errors.Is(err, ErrUserAlreadyExists) {
		fmt.Println("User already exists")
	}
}

func CreateService(repo UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (r MemoryUserRepository) Save(name string) error {
	return ErrUserAlreadyExists
}

func (s UserService) CreateUser(name string) error {
	if err := s.repo.Save(name); err != nil {
		return fmt.Errorf("create user %q: %w", name, err)
	}

	return nil
}
