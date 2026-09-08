package user

type Repository interface {
	Save(name string) error
}

type InMemoryRepository struct {
	users map[string]bool
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users: make(map[string]bool),
	}
}

func (r *InMemoryRepository) Save(name string) error {
	if r.users[name] {
		return ErrUserExists
	}

	r.users[name] = true
	return nil
}
