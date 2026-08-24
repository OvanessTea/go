package user

type User struct {
	Name string
	Age  int
}

type Notifier interface {
	Notify(message string)
}

type Service struct {
	Notifier
}

func NewUser(name string, age int) User {
	return User{
		Name: name,
		Age:  age,
	}
}

func NewService(n Notifier) Service {
	return Service{
		Notifier: n,
	}
}

func (s Service) SendNotification(message string) {
	s.Notify(message)
}
