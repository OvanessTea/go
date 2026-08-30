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
		{Name: "Kate", Age: 31},
	}

	for _, user := range users {
		fmt.Println(user.Name)
	}

	users = append(users, User{
		Name: "Mike",
		Age:  25,
	})

	fmt.Println(len(users))

	user, err := findUser(users, "Bob")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)

	user, err = findUser(users, "John")

	if err != nil {
		fmt.Println(err)
		// return commented for next task
	}
	fmt.Println(user)

	users = []User{
		{Name: "Alex", Age: 35},
		{Name: "Bob", Age: 28},
		{Name: "Kate", Age: 17},
		{Name: "Mike", Age: 25},
	}

	adults := getAdults(users)
	fmt.Println(adults)
}

func findUser(users []User, name string) (User, error) {
	for _, user := range users {
		if user.Name == name {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User %v not found", name)
}

func getAdults(users []User) []User {
	var adults []User

	for _, user := range users {
		if user.Age >= 18 {
			adults = append(adults, user)
		}
	}

	return adults
}
