package main

import (
	"errors"
	"fmt"
)

var ErrUserExists = errors.New("user already exists")

type UserRepository interface {
	Save(name string) error
}

type MemoryUserRepository struct{}
type FailingUserRepository struct{}

type UserService struct {
	repo UserRepository
}

func main() {
	repo := MemoryUserRepository{}
	service := NewUserService(repo)

	err := service.CreateUser("Alex")

	if errors.Is(err, ErrUserExists) {
		fmt.Println("user already exists")
		return
	}
	fmt.Println("User successfully created")

	repoFail := FailingUserRepository{}
	service = NewUserService(repoFail)

	err = service.CreateUser("Alex")

	if errors.Is(err, ErrUserExists) {
		fmt.Println("user already exists")
		return
	}
	fmt.Println("User successfully created")
}

func (r MemoryUserRepository) Save(name string) error {
	return nil
}

func (r FailingUserRepository) Save(name string) error {
	return ErrUserExists
}

func NewUserService(repo UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) CreateUser(name string) error {
	if err := s.repo.Save(name); err != nil {
		return fmt.Errorf("create user %q: %w", name, err)
	}

	return nil
}

// Вопрос 1
// Потому что добавление контекста через %w позволяет оборачивать именно ошибку
// , когда как %v используется для обычных строк

// Вопрос 2
// Будет false. Потому что это уже не будет ошибкой конкретного типа

// Вопрос 3
// Handler. Потому что иные слои получают, но не обрабатывают ошибки. Они их просто поднимают на уровень выше.
// А Handler уже решает, что делать с той или иной ошибкой

// Вопрос 4
// Потому что в первом случае мы привязываем сервис к конкретной реализациИ
// , что ограничивает приложение к расширению. Когда как зависимость от интерфейса
// позволяет нам добавлять новые реализации и структуры без необходимости менять сервис
