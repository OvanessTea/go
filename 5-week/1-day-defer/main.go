package main

import "fmt"

// panic - is used when program couldn't complete execution
// panic != error
func main() {
	// test()
	// code stops to execute after triggering panic
	// fmt.Println("hello")
	//fmt won't trigger

	recovering()

	fmt.Println("main continues") // recover allows to continue execution
}

func test() {
	defer fmt.Println("cleanup")
	// panic doesn't cancel defer
	panic("boom")
}

func recovering() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovored:", r) // recovored: boom
		}
	}() // recover has to be located in defer. 'cause recover needs to be exe while panic

	panic("boom")

	fmt.Println("never executed") // never executed
}
