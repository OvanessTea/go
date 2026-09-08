package user

type Repository interface {
	Save(name string) error
	FindAll() []string
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

func (r *InMemoryRepository) FindAll() []string {
	users := make([]string, 0, len(r.users))

	for name := range r.users {
		users = append(users, name)
	}

	return users
}
