package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(result)

	// result, err = divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(result)

	user, err := findUser(5)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(user)
}

// Task 1
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

// Task 2
type User struct {
	Name string
	Age  int
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
