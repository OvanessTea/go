package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

type ValidationError struct {
	Field string
}

type User struct {
	Name string
	Age  int
}

func main() {
	user, err := getUser(42)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			fmt.Println(err)
			fmt.Println("User does not exist")
			// return - comment to prevent program stop
		}
	}

	fmt.Println(user)

	var validationErr ValidationError

	err = validateUser(user)

	if errors.As(err, &validationErr) {
		fmt.Println("Invalid field:", validationErr.Field)
	}
}

func findUser(id int) (User, error) {
	if id != 1 {
		return User{}, ErrUserNotFound
	}

	return User{"Alex", 35}, nil
}

func getUser(id int) (User, error) {
	user, err := findUser(id)

	if err != nil {
		return User{}, fmt.Errorf("failed to get user %d: %w", id, err)
	}

	return user, nil
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("invalid field: %s", e.Field)
}

func validateUser(user User) error {
	if user.Name == "" {
		return ValidationError{Field: "name"}
	}
	if user.Age < 0 {
		return ValidationError{Field: "age"}
	}

	return nil
}
