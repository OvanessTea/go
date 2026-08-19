package main

import (
	"fmt"
)

type Speaker interface {
	Speak()
}
type Dog struct {
	Name string
}
type Cat struct {
	Name string
}

// Task 4
type Address struct {
	City   string
	Street string
}
type User struct {
	Name string
	Address
}

func main() {
	// Task 2 & 3
	makeSpeak(&Dog{Name: "Rex"}) // works
	// makeSpeak(Dog{Name: "Rex"}) // Dog does not implement Speaker (method Speak has pointer receiver)
	// 'cause we used pointer receiver in method
	// 	Dog
	//  ↓
	// methods with receiver (Dog)

	// *Dog
	//	↓
	// methods with receiver (Dog)
	// +
	// methods with receiver (*Dog)

	makeSpeak(Cat{Name: "Milo"})

	// Task 4
	user := User{
		Name: "Alice",
		Address: Address{
			City:   "Tokyo",
			Street: "Ahegao",
		},
	}
	fmt.Println(user.City)
	fmt.Println(user.Street)

	// Task 5
	inspect("hello")
	inspect(42)
	inspect(true)
	inspect(3.14)

	// Task to think 1
	// speaker.Speak() -> Woof
	// 'cause:
	// 1. Dog has method Speak()
	// 2. All structs what have Speak() implement Speaker interface
	// 3. var speaker has type Speaker
	// 4. After initialization of Dog{} it has dynamic type Dog
	// 5. It could use any Dog methods

	// Task to think 2
	// error 'cause we use pointer receiver
	// to fix: use &Dog{} instead of Dog{}

	// Task to think 3
	// str -> zero value string
	// ok -> false
	// 'cause str waits string type value but it takes int

	// Task to think 4
	// Panic
	// 'cause we don't use comm-ok

	// Task to think 5
	// false
	// 'cause speaker is not a nil
	// it has dynamic type Dog and dynamic value - nil
	// Go sees dynamic type != nil and tells speaker != nil

}

// Task 1
func (d *Dog) Speak() {
	fmt.Println("Wook")
}
func (c Cat) Speak() {
	fmt.Println("Meow")
}

// Task 2
func makeSpeak(s Speaker) {
	s.Speak()
}

// Task 5
func inspect(value any) {
	switch value.(type) {
	case string:
		fmt.Println("String:", value)
	case int:
		fmt.Println("Int:", value)
	case bool:
		fmt.Println("Bool:", value)
	default:
		fmt.Println("unknow type")
	}
}
