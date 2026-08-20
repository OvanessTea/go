package main

// Package - a way to combine chained Go-code and rule its accession

import (
	"day8/user"
	"fmt"
)

func main() {
	u := user.User{
		Name: "Alex",
	}

	fmt.Println(u.Name)
}

//          Go module
//             │
//       ┌─────┴─────┐
//       │            │
//    package A    package B
//       │            │
// exported       exported
// unexported     unexported
