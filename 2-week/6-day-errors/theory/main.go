package main

import (
	"errors"
	"fmt"
)

// Default Go Error interface
// type error interface {
// 	Error() string
// }

func main() {
	result, err := divide(10, 2)

	user, err := findUser(2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(result)
	fmt.Println(user)

}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return a / b, nil
}

func findUser(id int) (string, error) {
	if id != 1 {
		return "", fmt.Errorf("user %d not found", id) // fmt.Errorf to include data
	}

	return "Alex", nil
}

// Error is a value.
// err := errors.New("something went wrong")
// err - usual variable what has value. That value satisfies interface error
