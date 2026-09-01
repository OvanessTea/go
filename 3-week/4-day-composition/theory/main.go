package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Address // embedded field
}

type Address struct {
	City   string
	Street string
}

func main() {
	user, err := NewUser("Alex", -10)

	if err != nil {
		fmt.Println(err)
		// return
	}

	// fmt.Println(user)

	// composition
	user = User{
		Name: "Alex",
		Age:  35,
		Address: Address{
			City:   "Berlin",
			Street: "Main Street",
		},
	}

	fmt.Println(user.City) // promoted field

	fmt.Println(user.FullAddress()) // Methords are also promoted
	// 	User
	//  └── Address
	//        └── FullAddress()

	// It's not a inheritance
	// ||| Composition over inheritance |||
}

// constructor allows to add validation
// constructor is not required
func NewUser(name string, age int) (User, error) { // New... name is an agreement in Go
	if age < 0 {
		return User{}, fmt.Errorf("invalid age: %d", age)
	}

	return User{
		Name:    name,
		Age:     age,
		Address: Address{},
	}, nil
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s - %s", a.City, a.Street)
}
