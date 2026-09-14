package main

import "fmt"

func main() {
	hello()
}

// defer is used when we want to call smth at the end of function
// defers execute in reversed order: LIFO: last in, first out.
func hello() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("start")
}
