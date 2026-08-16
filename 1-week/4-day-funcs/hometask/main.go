package main

import (
	"errors"
	"fmt"
)

func main() {
	user, err := findUser(2)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User:", user)
}

func findUser(id int64) (string, error) {
	users := map[int64]string{
		1: "Alice",
		2: "Bob",
		3: "Charlie",
	}
	for key, name := range users {
		if id == key {
			return name, nil
		}
	}
	return "", errors.New("User not found")
}
