package main

import "fmt"

type User struct {
	ID       int64
	Name     string
	Age      int
	IsActive bool
}

func main() {
	users := []User{
		{ID: 1, Name: "Alice", Age: 25, IsActive: true},
		{ID: 2, Name: "Bob", Age: 17, IsActive: false},
		{ID: 3, Name: "Charlie", Age: 30, IsActive: true},
	}
	ok := activateUser(users, 1)
	if !ok {
		fmt.Println("User not found")
		return
	}
	fmt.Println("User found and changed its activity")
	fmt.Println(users)
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (u *User) Active() {
	u.IsActive = true
}

func (u *User) Deactive() {
	u.IsActive = false
}

func (u *User) ChangeName(name string) {
	u.Name = name
}

func findUser(users []User, id int64) (*User, bool) {
	for i := range users {
		if users[i].ID == id {
			return &users[i], true
		}
	}
	var nilUser User
	return &nilUser, false
}

func activateUser(users []User, id int64) bool {
	user, ok := findUser(users, id)

	if !ok {
		return false
	}

	if user.IsActive {
		user.Deactive()
	} else {
		user.Active()
	}
	return true
}
