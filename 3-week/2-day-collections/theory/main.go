package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{Name: "Alex", Age: 35},
		{Name: "Bob", Age: 28},
	} // type - []User

	users = append(users, User{
		Name: "Kate",
		Age:  30,
	})

	fmt.Println(users)
	fmt.Println(len(users))
}
