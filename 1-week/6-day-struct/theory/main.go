package main

import "fmt"

type User struct {
	ID   int64
	Name string
	Age  int
}

func main() {
	// STRUCT
	user := User{
		ID:   1,
		Name: "Alice",
		Age:  25,
	}
	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	user.Age = 26
	fmt.Println(user.Age)

	var zeroUser User // zero value for struct
	fmt.Println(zeroUser)

	printUser(user)
	fmt.Println(user.Age) // source keeps its value

	// POINTER
	ptr := &user // pointer - address of the value
	// var ptr *User // ptr is a pointer to User
	fmt.Printf("%T\n", ptr) // *User type
	fmt.Println(ptr)        // address
	fmt.Println(*ptr)       // dereference - get a value from that pointer

	ptr.Age = 100 // Allows to change a value via pointer
	// ptr.Age == (*ptr).Age - in Go it's automatic
	fmt.Println(user.Age)

	changeAge(&user)  // Updates original struct 100 -> 200
	fmt.Println(user) // 200

	var nilUser *User // zero value pointer
	fmt.Println(nilUser)
	// fmt.Println(nilUser.Name) -> will trigger Panic

	// STRUCT METHODS
	// We can create a method for existed struct
	fmt.Println(user.IsAdult())
	user.ChangeName("Bob") // Go use address automatically
	fmt.Println(user.Name) //Bob

	// STRUCT + SLICE
	users := []User{
		{ID: 1, Name: "Alice", Age: 25},
		{ID: 2, Name: "Bob", Age: 30},
	}

	for _, user := range users {
		user.Age++
	}
	fmt.Println(users) // ages are the same 'cause we 're not using pointer
	// So inside the loop we have the COPIES of structs

	// To update original structs we have to use for i
	for i := range users {
		users[i].Age++ // this will work 'cause slice[i] direct to the original slice elem
	}
	fmt.Println(users)
}

// inside function we have another indipendend "copy" of proped struct
func printUser(user User) { //struct as an argument
	fmt.Println(user.Name)
	user.Age = 50 // it doesn't change a source struct
	fmt.Println(user.Age)
}

func changeAge(user *User) { // 'cause we use Pointer as prop, we can change the source struct
	user.Age = 200
}

func (user User) IsAdult() bool { // (user User) - receiver
	return user.Age >= 18
}

// receiver shows what struct includes current method
// User
//  ↓
// IsAdult()

func (user *User) ChangeName(name string) { // 'cause we use pointer we can change the source value via method
	user.Name = name
}

// We don't need to use Pointer receiver everytime.
// If we don't need to change original struct, we should not use pointer receiver
