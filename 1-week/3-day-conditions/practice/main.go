package main

import "fmt"

func main() {
	// Task 1
	age := 25

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// Task 2
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}

	// Task 3
	counter := 10

	for counter >= 0 {
		fmt.Println(counter)

		counter--
	}

	// Task 4
	status := "pending"

	switch status {
	case "pending":
		fmt.Println("Waiting")
	case "active":
		fmt.Println("Running")
	case "completed":
		fmt.Println("Done")
	case "failed":
		fmt.Println("Failed")
	default:
		fmt.Println("Unknown")
	}

	// Task 5
	names := []string{"Alice", "Bob", "Charlie", "Dave"}

	for _, name := range names {
		fmt.Println("Hello, " + name)
	}

	// Task to think
	// 0, 1, 3, 4

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}

		fmt.Println(i)
	}
}
