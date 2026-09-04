package main

import "fmt"

type UserRepository interface {
	Save(name string)
}

type MockUserRepository struct {
	SavedName string
}

func (r *MockUserRepository) Save(name string) {
	r.SavedName = name
}

func main() {
	mock := MockUserRepository{}
	service := CreateUserService(&mock)

	service.CreateUser("Alex")

	fmt.Println(mock.SavedName)
}

type UserService struct {
	repo UserRepository
}

func CreateUserService(repo UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) CreateUser(name string) {
	s.repo.Save(name)
}

// Вопрос: Почему здесь MockUserRepository имеет pointer receiver (*MockUserRepository), а не value receiver?
// Потому что мы хотим изменить оригинал, а не копию структуры. Поэтому мы обращаемся по адресу для изменения параметра
