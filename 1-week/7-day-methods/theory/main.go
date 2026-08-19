package main

import (
	"fmt"
)

// struct - for data
// interface - for methods
// interface could include several methods
type Animal interface {
	Speak()
	Move()
}

type Stopper interface {
	Stop()
}

type Dog struct{}
type Cat struct{}

// Composition
type Address struct {
	City   string
	Street string
}
type User struct {
	Name    string
	Address // Embedded struct - don't need to write Address Address
}

// User
// ├── Name
// └── Address
//     ├── City
//     └── Street

// Embedding != inheritance
// User IS-A Address - WRONG
// User HAS-A Address - CORRECT

// EMBEDDING of METHODS
type Logger struct{}

func (l Logger) Log(message string) {
	fmt.Println(message)
}

type Service struct {
	Logger
}

// Composition + interface
type ILogger interface {
	Log(string)
}
type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

func process(logger ILogger) {
	logger.Log("processing")
} // possible 'cause that method need only contruct Log(string)

// ANY
// any is an alias for '''interface{}'''
// var value any - could has ANY type value inside
// 'cause any type satisfy an empty interface

func main() {
	// Interface
	dog := Dog{}
	cat := Cat{}

	makeSpeak(dog)
	makeSpeak(cat)

	var animal Animal //variable could includes any object with that interface
	animal = Dog{}
	animal.Move()

	animal = Cat{}
	animal.Move()

	// Dynamic type & dynamic value
	// var animal Animal = Dog{}
	// 	interface
	// ├── dynamic type  → Dog
	// └── dynamic value → Dog{}

	// animal = Cat{}
	// 	interface
	// ├── dynamic type  → Cat
	// └── dynamic value → Cat{}

	// But type of variable is static -> Animal

	// Nil value interface
	var emptyAnimal Animal
	fmt.Println(emptyAnimal == nil) // true
	// type - nil | value - nil
	// But
	var nilDog *Dog = nil
	emptyAnimal = nilDog
	fmt.Println(emptyAnimal == nil) // false -> 'cause it has dynamic type *Dog even though dynamic value == nil
	// type - *Dog | value - nil

	// Pointer receiver & interface
	var stopper Stopper
	stopper = &Dog{} // works
	// stopper = Dog{} // does not work 'cause we use pointer receiver in method
	stopper.Stop()
	// Dog & *Dog have different method sets

	// 	Dog
	//  ↓
	// methods with receiver (Dog)

	// *Dog
	//	↓
	// methods with receiver (Dog)
	// +
	// methods with receiver (*Dog)

	// COMPOSITION + EMBEDDED STRUCT
	user := User{
		Name: "Alice",
		Address: Address{
			City:   "Amsterdam",
			Street: "Main",
		},
	}

	fmt.Println(user.City) // no need to write user.Address.City
	// Fields of embedded struct promoted in enclosing struct

	// EMBEDDING of METHODS
	service := Service{}
	service.Log("HELLO") // Method promoted from Logger struct

	// TYPE ASSERTION
	var value any = "hello" // we know that's string
	str := value.(string)   // str has type string
	fmt.Printf("%T\n", str)
	// value = 123
	// str := value.(string) // will trigger Panic 'cause str waits string value

	// Safe type assertion
	str, ok := value.(string) // comma-ok

	if !ok {
		fmt.Println("not a string")
		return
	}
	fmt.Println(str)

	// TYPE SWITCH
	switch value := value.(type) {
	case string:
		fmt.Println("string:", value)
	case int:
		fmt.Println("int:", value)
	case bool:
		fmt.Println("bool:", value)
	default:
		fmt.Println("unknown")
	}
}

// If struct uses the same methods as interface -> struct implements that interface
// Implicit interface implemention
// To implement interface struct has to have all methods of that interface
func (d Dog) Speak() {
	fmt.Println("Woof")
}

func (d Dog) Move() {
	fmt.Println("Running")
}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

func (c Cat) Move() {
	fmt.Println("Moving")
}

// Dog ──┐
//       ├── Speaker
// Cat ──┘
// 'cause them use the same methods

func makeSpeak(s Animal) {
	s.Speak()
}

func (d *Dog) Stop() { // pointer receiver
	fmt.Println("Stopped")
}
