package main

import "fmt"

type User struct {
	ID   int64
	Name string
	Age  int
}

func main() {
	users := []User{
		{ID: 1, Name: "Alice", Age: 25},
		{ID: 2, Name: "Bob", Age: 30},
		{ID: 3, Name: "Charlie", Age: 35},
	}

	fmt.Println(findUser(users, 1))
	fmt.Println(findUser(users, 999))

	users = addUser(users, User{ID: 4, Name: "Dave", Age: 50})
	fmt.Println(findUser(users, 4))

	users = deleteUser(users, 2)

	user, ok := findUserByName(users, "Bob")
	if !ok {
		fmt.Println("User not found")
		return
	}
	fmt.Println(user)
}

func findUser(users []User, id int64) (User, bool) {
	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}
	var emptyUser User
	return emptyUser, false
}

func addUser(users []User, user User) []User {
	return append(users, user)
}

func deleteUser(users []User, id int64) []User {
	for index, user := range users {
		if user.ID == id {
			return append(users[:index], users[index+1:]...)
		}
	}
	return users
}

func findUserByName(users []User, name string) (User, bool) {
	for _, user := range users {
		if user.Name == name {
			return user, true
		}
	}
	var emptyUser User
	return emptyUser, false
}
