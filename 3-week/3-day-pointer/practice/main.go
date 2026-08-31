package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Describer interface {
	Describe() string
}

func main() {
	u := User{Name: "Alex", Age: 25}

	fmt.Println(u.IsAdult())

	u.SetAge(40)

	fmt.Println(u.Age) // результат меняется, если мы используем pointer receiver
	// в таком случае меняется не копия структуры, а оригинал

	var d Describer = &u
	// до исправления был Panic:
	// .\main.go:24:20: cannot use u (variable of struct type User) as Describer value in variable declaration: User does not implement Describer (method Describe has pointer receiver)
	fmt.Println(d.Describe())
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func (u *User) Describe() string {
	return fmt.Sprintf("%s, %d years old", u.Name, u.Age)
}
