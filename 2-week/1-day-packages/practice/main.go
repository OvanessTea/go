package main

import (
	"fmt"
	"practice/user"
)

func main() {
	u := user.User{
		Name: "Alex",
		Age:  25,
		// Task 2
		// email: "Aboba@example.com",
		// cannot refer to unexported field email in struct literal of type
	}

	fmt.Println(u.Name)

	// Task 3
	newU := user.NewUser("Alice", 30)

	fmt.Println(newU)

	// Task 4
	// It allows to hide some params or methods 'cause they should be used only in package where they were created
	// Like "private" params and methods
}
