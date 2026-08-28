package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user, err := getUser(2)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
}

func getUser(id int) (User, error) {
	user, err := findUser(id)

	if err != nil {
		return User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func findUser(id int) (User, error) {
	if id != 1 {
		return User{}, fmt.Errorf("used %d not found", id)
	}
	return User{"Alex", 25}, nil
}
