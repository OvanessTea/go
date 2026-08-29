package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	Name string
	Age  int
}

type ValidationError struct {
	Field string
}

func main() {
	// user, err := findUser(42)

	// if errors.Is(err, ErrUserNotFound) {
	// 	fmt.Println("User does not exist")
	// 	return
	// }

	// fmt.Println(user)
	var validationErr ValidationError

	err := ValidationError{
		Field: "email",
	}

	if errors.As(err, &validationErr) {
		fmt.Println("Ivalid field:", validationErr.Field)
	}
}

func findUser(id int) (User, error) {
	if id != 1 {
		return User{}, fmt.Errorf("failted to get user: %w", ErrUserNotFound)
	}

	return User{"Alex", 35}, nil
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("invalid field: %s", e.Field)
}
