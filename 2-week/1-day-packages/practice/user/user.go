package user

// Task 1
type User struct {
	Name string
	Age  int
	// Task 2
	email string
}

// Task 3
func NewUser(name string, age int) User {
	u := User{Name: name, Age: age}
	return u
}
