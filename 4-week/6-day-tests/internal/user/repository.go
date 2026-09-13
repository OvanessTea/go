package user

import (
	"myapp/internal/dto"
)

type Repository interface {
	Save(name string) error
	FindAll() []dto.User
	FindByID(id int) (dto.User, error)
}

type InMemoryRepository struct {
	users  map[int]dto.User
	nextID int
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users: make(map[int]dto.User),
	}
}

func (r *InMemoryRepository) Save(name string) error {

	for _, user := range r.users {
		if user.Name == name {
			return ErrUserExists
		}
	}

	r.users[r.nextID] = dto.User{
		ID:   r.nextID,
		Name: name,
	}
	r.nextID++
	return nil
}

func (r *InMemoryRepository) FindAll() []dto.User {
	users := make([]dto.User, 0, len(r.users))

	for _, user := range r.users {
		users = append(users, user)
	}

	return users
}

func (r *InMemoryRepository) FindByID(id int) (dto.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return dto.User{}, ErrUserNotFound
}
