package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user, err := getUser(42)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(user)
}

func findUser(id int) (User, error) {
	switch id {
	case 1:
		return User{"Alex", 35}, nil
	case 2:
		return User{"Bob", 28}, nil
	}
	return User{}, fmt.Errorf("user %d not found", id)
}

func getUser(id int) (User, error) {
	user, err := findUser(id)

	if err != nil {
		return User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}
